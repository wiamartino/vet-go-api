package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AllergyController struct {
	service *application.AllergyService
}

func NewAllergyController(service *application.AllergyService) *AllergyController {
	return &AllergyController{service: service}
}

// GetAllergies retrieves all allergies
// @Summary Get all allergies
// @Description Get a list of all recorded allergies
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Allergy "List of allergies"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /allergies [get]
func (c *AllergyController) GetAllergies(ctx *gin.Context) {
	allergies, err := c.service.GetAllAllergies()
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}
	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

// GetAllergy retrieves an allergy by ID
// @Summary Get allergy by ID
// @Description Get detailed information about a specific allergy
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Allergy ID"
// @Success 200 {object} domain.Allergy "Allergy details"
// @Failure 400 {object} map[string]string "Invalid allergy ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Allergy not found"
// @Router /allergies/{id} [get]
func (c *AllergyController) GetAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid allergy ID"))
		return
	}

	allergy, err := c.service.GetAllergyByID(uint(id))
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Allergy"))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergy)
}

// GetAllergiesByPet retrieves allergies for a specific pet
// @Summary Get allergies by pet ID
// @Description Get all allergies for a specific pet
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {array} domain.Allergy "List of allergies for the pet"
// @Failure 400 {object} map[string]string "Invalid pet ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets/{id}/allergies [get]
func (c *AllergyController) GetAllergiesByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid pet ID"))
		return
	}

	allergies, err := c.service.GetAllergiesByPetID(uint(petID))
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

// GetActiveAllergiesByPet retrieves active allergies for a specific pet
// @Summary Get active allergies by pet ID
// @Description Get all active allergies for a specific pet
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {array} domain.Allergy "List of active allergies for the pet"
// @Failure 400 {object} map[string]string "Invalid pet ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets/{id}/allergies/active [get]
func (c *AllergyController) GetActiveAllergiesByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid pet ID"))
		return
	}

	allergies, err := c.service.GetActiveAllergiesByPetID(uint(petID))
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

// GetAllergiesBySeverity retrieves allergies by severity
// @Summary Get allergies by severity
// @Description Get all allergies filtered by severity level
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Param severity query string true "Allergy severity (mild, moderate, severe, fatal)"
// @Success 200 {array} domain.Allergy "List of allergies with the specified severity"
// @Failure 400 {object} map[string]string "Severity parameter is required"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /allergies/severity [get]
func (c *AllergyController) GetAllergiesBySeverity(ctx *gin.Context) {
	severityStr := ctx.Query("severity")
	if severityStr == "" {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Severity parameter is required"))
		return
	}

	severity := domain.AllergySeverity(severityStr)
	allergies, err := c.service.GetAllergiesBySeverity(severity)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

// CreateAllergy creates a new allergy
// @Summary Create a new allergy
// @Description Record a new allergy for a pet
// @Tags Allergies
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param allergy body domain.Allergy true "Allergy details"
// @Success 201 {object} domain.Allergy "Allergy created successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /allergies [post]
func (c *AllergyController) CreateAllergy(ctx *gin.Context) {
	var allergy domain.Allergy
	if err := ctx.ShouldBindJSON(&allergy); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	validator := utils.NewValidator()
	validator.ValidateNumericID("pet_id", allergy.PetID)
	validator.ValidateRequired("allergen", allergy.Allergen)
	validator.ValidateInSlice("allergy_type", string(allergy.AllergyType), []string{"food", "medication", "environment", "insect", "other"})
	validator.ValidateInSlice("severity", string(allergy.Severity), []string{"mild", "moderate", "severe", "fatal"})

	if !validator.IsValid() {
		utils.RespondWithValidationError(ctx, validator.Errors)
		return
	}

	if err := c.service.CreateAllergy(&allergy); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithCreated(ctx, allergy)
}

// UpdateAllergy updates an allergy by ID
// @Summary Update allergy by ID
// @Description Update information for a specific allergy
// @Tags Allergies
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Allergy ID"
// @Param allergy body domain.Allergy true "Allergy details"
// @Success 200 {object} domain.Allergy "Allergy updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /allergies/{id} [put]
func (c *AllergyController) UpdateAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid allergy ID"))
		return
	}

	// Check if allergy exists
	if _, err := c.service.GetAllergyByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Allergy"))
		return
	}

	var allergy domain.Allergy
	if err := ctx.ShouldBindJSON(&allergy); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	validator := utils.NewValidator()
	validator.ValidateNumericID("pet_id", allergy.PetID)
	validator.ValidateRequired("allergen", allergy.Allergen)
	validator.ValidateInSlice("allergy_type", string(allergy.AllergyType), []string{"food", "medication", "environment", "insect", "other"})
	validator.ValidateInSlice("severity", string(allergy.Severity), []string{"mild", "moderate", "severe", "fatal"})

	if !validator.IsValid() {
		utils.RespondWithValidationError(ctx, validator.Errors)
		return
	}

	allergy.AllergyID = uint(id)
	if err := c.service.UpdateAllergy(&allergy); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergy)
}

// DeactivateAllergy deactivates an allergy
// @Summary Deactivate an allergy
// @Description Mark an allergy as inactive
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Allergy ID"
// @Success 200 {object} map[string]string "Allergy deactivated successfully"
// @Failure 400 {object} map[string]string "Invalid allergy ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /allergies/{id}/deactivate [patch]
func (c *AllergyController) DeactivateAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid allergy ID"))
		return
	}

	if _, err := c.service.GetAllergyByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Allergy"))
		return
	}

	if err := c.service.DeactivateAllergy(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Allergy deactivated successfully"})
}

// ReactivateAllergy reactivates an allergy
// @Summary Reactivate an allergy
// @Description Mark an allergy as active again
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Allergy ID"
// @Success 200 {object} map[string]string "Allergy reactivated successfully"
// @Failure 400 {object} map[string]string "Invalid allergy ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /allergies/{id}/reactivate [patch]
func (c *AllergyController) ReactivateAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid allergy ID"))
		return
	}

	if _, err := c.service.GetAllergyByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Allergy"))
		return
	}

	if err := c.service.ReactivateAllergy(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Allergy reactivated successfully"})
}

// DeleteAllergy deletes an allergy by ID
// @Summary Delete allergy by ID
// @Description Remove an allergy record from the system
// @Tags Allergies
// @Security BearerAuth
// @Produce json
// @Param id path int true "Allergy ID"
// @Success 204 "Allergy deleted successfully"
// @Failure 400 {object} map[string]string "Invalid allergy ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /allergies/{id} [delete]
func (c *AllergyController) DeleteAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid allergy ID"))
		return
	}

	if _, err := c.service.GetAllergyByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Allergy"))
		return
	}

	if err := c.service.DeleteAllergy(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithNoContent(ctx)
}
