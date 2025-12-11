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

func (ctrl *PetController) FindPets(c *gin.Context) {
	pets, err := ctrl.service.GetAllPets()
	if err != nil {
		utils.RespondWithAppError(c, utils.AsAppError(err))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, pets)
}

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
