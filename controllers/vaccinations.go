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

// GetVaccinations retrieves all vaccinations
// @Summary Get all vaccinations
// @Description Get a list of all vaccinations
// @Tags Vaccinations
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Vaccination "List of vaccinations"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /vaccinations [get]
func (c *VaccinationController) GetVaccinations(ctx *gin.Context) {
	vaccinations, err := c.service.GetAllVaccinations()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(ctx, http.StatusOK, vaccinations)
}

// GetVaccination retrieves a vaccination by ID
// @Summary Get vaccination by ID
// @Description Get detailed information about a specific vaccination
// @Tags Vaccinations
// @Security BearerAuth
// @Produce json
// @Param id path int true "Vaccination ID"
// @Success 200 {object} domain.Vaccination "Vaccination details"
// @Failure 400 {object} map[string]string "Invalid vaccination ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Vaccination not found"
// @Router /vaccinations/{id} [get]
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

// GetVaccinationsByPet retrieves vaccinations for a specific pet
// @Summary Get vaccinations by pet ID
// @Description Get all vaccinations for a specific pet
// @Tags Vaccinations
// @Security BearerAuth
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {array} domain.Vaccination "List of vaccinations for the pet"
// @Failure 400 {object} map[string]string "Invalid pet ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets/{id}/vaccinations [get]
func (c *VaccinationController) GetVaccinationsByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
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

// GetDueVaccinations retrieves vaccinations that are due
// @Summary Get due vaccinations
// @Description Get all vaccinations that are due within a specified number of days
// @Tags Vaccinations
// @Security BearerAuth
// @Produce json
// @Param days query int false "Number of days to look ahead (default 30)"
// @Success 200 {array} domain.Vaccination "List of due vaccinations"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /vaccinations/due [get]
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

// GetOverdueVaccinations retrieves overdue vaccinations
// @Summary Get overdue vaccinations
// @Description Get all vaccinations that are overdue
// @Tags Vaccinations
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Vaccination "List of overdue vaccinations"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /vaccinations/overdue [get]
func (c *VaccinationController) GetOverdueVaccinations(ctx *gin.Context) {
	vaccinations, err := c.service.GetOverdueVaccinations()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, vaccinations)
}

// CreateVaccination creates a new vaccination
// @Summary Create a new vaccination
// @Description Schedule a new vaccination for a pet
// @Tags Vaccinations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param vaccination body domain.Vaccination true "Vaccination details"
// @Success 201 {object} domain.Vaccination "Vaccination created successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /vaccinations [post]
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

// UpdateVaccination updates a vaccination by ID
// @Summary Update vaccination by ID
// @Description Update information for a specific vaccination
// @Tags Vaccinations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Vaccination ID"
// @Param vaccination body domain.Vaccination true "Vaccination details"
// @Success 200 {object} domain.Vaccination "Vaccination updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /vaccinations/{id} [put]
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

// CompleteVaccination marks a vaccination as completed
// @Summary Complete a vaccination
// @Description Mark a vaccination as completed with veterinarian details
// @Tags Vaccinations
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Vaccination ID"
// @Param body body object{veterinarian_id=int,side_effects=string} true "Completion details"
// @Success 200 {object} map[string]string "Vaccination completed successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /vaccinations/{id}/complete [post]
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

// DeleteVaccination deletes a vaccination by ID
// @Summary Delete vaccination by ID
// @Description Remove a vaccination from the system
// @Tags Vaccinations
// @Security BearerAuth
// @Produce json
// @Param id path int true "Vaccination ID"
// @Success 200 {object} map[string]string "Vaccination deleted successfully"
// @Failure 400 {object} map[string]string "Invalid vaccination ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /vaccinations/{id} [delete]
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
