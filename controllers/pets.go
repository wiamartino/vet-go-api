package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PetController struct {
	service *application.PetService
}

func NewPetController(service *application.PetService) *PetController {
	return &PetController{service: service}
}

// FindPets retrieves all pets
// @Summary Get all pets
// @Description Get a list of all registered pets
// @Tags Pets
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Pet "List of pets"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets [get]
func (ctrl *PetController) FindPets(c *gin.Context) {
	pets, err := ctrl.service.GetAllPets()
	if err != nil {
		utils.RespondWithAppError(c, utils.AsAppError(err))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, pets)
}

// CreatePet creates a new pet
// @Summary Create a new pet
// @Description Register a new pet in the system
// @Tags Pets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param pet body domain.Pet true "Pet details"
// @Success 201 {object} domain.Pet "Pet created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or validation errors"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets [post]
func (ctrl *PetController) CreatePet(c *gin.Context) {
	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	// Validate pet data
	validator := utils.NewValidator()
	validator.ValidateLengthRange("name", pet.Name, 2, 100)
	validator.ValidateLengthRange("species", pet.Species, 2, 50)
	validator.ValidateLengthRange("breed", pet.Breed, 2, 100)
	validator.ValidateNumericID("client_id", pet.ClientID)

	if !validator.IsValid() {
		utils.RespondWithValidationError(c, validator.Errors)
		return
	}

	if err := ctrl.service.CreatePet(&pet); err != nil {
		if utils.IsAppError(err) {
			utils.RespondWithAppError(c, err.(*utils.AppError))
		} else {
			utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		}
		return
	}
	utils.RespondWithCreated(c, pet)
}

// FindPet retrieves a pet by ID
// @Summary Get pet by ID
// @Description Get detailed information about a specific pet
// @Tags Pets
// @Security BearerAuth
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {object} domain.Pet "Pet details"
// @Failure 400 {object} map[string]string "Invalid pet ID format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Pet not found"
// @Router /pets/{id} [get]
func (ctrl *PetController) FindPet(c *gin.Context) {
	petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid pet ID format"))
		return
	}
	pet, err := ctrl.service.GetPetByID(uint(petID))
	if err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Pet"))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, pet)
}

// UpdatePet updates a pet by ID
// @Summary Update pet by ID
// @Description Update information for a specific pet
// @Tags Pets
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Pet ID"
// @Param pet body domain.Pet true "Pet details"
// @Success 200 {object} domain.Pet "Pet updated successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or validation errors"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Pet not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets/{id} [put]
func (ctrl *PetController) UpdatePet(c *gin.Context) {
	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid pet ID format"))
		return
	}

	// Check if pet exists
	if _, err := ctrl.service.GetPetByID(uint(petID)); err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Pet"))
		return
	}

	// Validate pet data
	validator := utils.NewValidator()
	validator.ValidateLengthRange("name", pet.Name, 2, 100)
	validator.ValidateLengthRange("species", pet.Species, 2, 50)
	validator.ValidateLengthRange("breed", pet.Breed, 2, 100)

	if !validator.IsValid() {
		utils.RespondWithValidationError(c, validator.Errors)
		return
	}

	pet.PetID = uint(petID)
	err = ctrl.service.UpdatePet(&pet)
	if err != nil {
		if utils.IsAppError(err) {
			utils.RespondWithAppError(c, err.(*utils.AppError))
		} else {
			utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		}
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, pet)
}

// DeletePet deletes a pet by ID
// @Summary Delete pet by ID
// @Description Delete a specific pet from the system
// @Tags Pets
// @Security BearerAuth
// @Produce json
// @Param id path int true "Pet ID"
// @Success 204 "Pet deleted successfully"
// @Failure 400 {object} map[string]string "Invalid pet ID format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Pet not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets/{id} [delete]
func (ctrl *PetController) DeletePet(c *gin.Context) {
	petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid pet ID format"))
		return
	}

	// Check if pet exists before attempting to delete
	if _, err := ctrl.service.GetPetByID(uint(petID)); err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Pet"))
		return
	}

	err = ctrl.service.DeletePet(uint(petID))
	if err != nil {
		utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		return
	}

	utils.RespondWithNoContent(c)
}
