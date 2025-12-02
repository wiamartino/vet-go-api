package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VaccinationController struct {
	service *application.VaccinationService
}

func NewVaccinationController(service *application.VaccinationService) *VaccinationController {
	return &VaccinationController{service: service}
}

func (c *VaccinationController) GetVaccinations(ctx *gin.Context) {
	vaccinations, err := c.service.GetAllVaccinations()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(ctx, http.StatusOK, vaccinations)
}

func (c *VaccinationController) GetVaccination(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	vaccination, err := c.service.GetVaccinationByID(uint(id))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, vaccination)
}

func (c *VaccinationController) GetVaccinationsByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("pet_id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	vaccinations, err := c.service.GetVaccinationsByPetID(uint(petID))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, vaccinations)
}

func (c *VaccinationController) GetDueVaccinations(ctx *gin.Context) {
	days := 30
	if daysParam := ctx.Query("days"); daysParam != "" {
		parsedDays, err := strconv.Atoi(daysParam)
		if err == nil && parsedDays > 0 {
			days = parsedDays
		}
	}

	vaccinations, err := c.service.GetDueVaccinations(days)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, vaccinations)
}

func (c *VaccinationController) GetOverdueVaccinations(ctx *gin.Context) {
	vaccinations, err := c.service.GetOverdueVaccinations()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, vaccinations)
}

func (c *VaccinationController) CreateVaccination(ctx *gin.Context) {
	var vaccination domain.Vaccination
	if err := ctx.ShouldBindJSON(&vaccination); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CreateVaccination(&vaccination); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusCreated, vaccination)
}

func (c *VaccinationController) UpdateVaccination(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var vaccination domain.Vaccination
	if err := ctx.ShouldBindJSON(&vaccination); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	vaccination.VaccinationID = uint(id)
	if err := c.service.UpdateVaccination(&vaccination); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, vaccination)
}

func (c *VaccinationController) CompleteVaccination(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var body struct {
		VeterinarianID uint   `json:"veterinarian_id" binding:"required"`
		SideEffects    string `json:"side_effects"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CompleteVaccination(uint(id), body.VeterinarianID, body.SideEffects); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Vaccination completed successfully"})
}

func (c *VaccinationController) DeleteVaccination(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.DeleteVaccination(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Vaccination deleted successfully"})
}
