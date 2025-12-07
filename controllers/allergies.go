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

func (c *AllergyController) GetAllergies(ctx *gin.Context) {
	allergies, err := c.service.GetAllAllergies()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

func (c *AllergyController) GetAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	allergy, err := c.service.GetAllergyByID(uint(id))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergy)
}

func (c *AllergyController) GetAllergiesByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	allergies, err := c.service.GetAllergiesByPetID(uint(petID))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

func (c *AllergyController) GetActiveAllergiesByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	allergies, err := c.service.GetActiveAllergiesByPetID(uint(petID))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

func (c *AllergyController) GetAllergiesBySeverity(ctx *gin.Context) {
	severityStr := ctx.Query("severity")
	if severityStr == "" {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Severity parameter is required")
		return
	}

	severity := domain.AllergySeverity(severityStr)
	allergies, err := c.service.GetAllergiesBySeverity(severity)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergies)
}

func (c *AllergyController) CreateAllergy(ctx *gin.Context) {
	var allergy domain.Allergy
	if err := ctx.ShouldBindJSON(&allergy); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CreateAllergy(&allergy); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusCreated, allergy)
}

func (c *AllergyController) UpdateAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var allergy domain.Allergy
	if err := ctx.ShouldBindJSON(&allergy); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	allergy.AllergyID = uint(id)
	if err := c.service.UpdateAllergy(&allergy); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, allergy)
}

func (c *AllergyController) DeactivateAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.DeactivateAllergy(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Allergy deactivated successfully"})
}

func (c *AllergyController) ReactivateAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.ReactivateAllergy(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Allergy reactivated successfully"})
}

func (c *AllergyController) DeleteAllergy(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.DeleteAllergy(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Allergy deleted successfully"})
}
