package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type TreatmentController struct {
	service *application.TreatmentService
}

func NewTreatmentController(service *application.TreatmentService) *TreatmentController {
	return &TreatmentController{service: service}
}

// FindTreatments retrieves all treatments
// @Summary Get all treatments
// @Description Get a list of all available treatments
// @Tags Treatments
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Treatment "List of treatments"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /treatments [get]
func (ctrl *TreatmentController) FindTreatments(c *gin.Context) {
	treatments, err := ctrl.service.GetAllTreatments()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, treatments)
}

// CreateTreatment creates a new treatment
// @Summary Create a new treatment
// @Description Add a new treatment to the system
// @Tags Treatments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param treatment body domain.Treatment true "Treatment details"
// @Success 200 {object} domain.Treatment "Treatment created successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /treatments [post]
func (ctrl *TreatmentController) CreateTreatment(c *gin.Context) {
	var treatment domain.Treatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctrl.service.CreateTreatment(&treatment); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, treatment)
}

// FindTreatment retrieves a treatment by ID
// @Summary Get treatment by ID
// @Description Get detailed information about a specific treatment
// @Tags Treatments
// @Security BearerAuth
// @Produce json
// @Param id path int true "Treatment ID"
// @Success 200 {object} domain.Treatment "Treatment details"
// @Failure 400 {object} map[string]string "Invalid treatment ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Treatment not found"
// @Router /treatments/{id} [get]
func (ctrl *TreatmentController) FindTreatment(c *gin.Context) {
	treatmentIDStr := c.Param("id")
	treatmentID, err := strconv.ParseUint(treatmentIDStr, 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid treatment ID")
		return
	}
	treatment, err := ctrl.service.GetTreatmentByID(uint(treatmentID))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Treatment not found")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, treatment)
}

// UpdateTreatment updates a treatment by ID
// @Summary Update treatment by ID
// @Description Update information for a specific treatment
// @Tags Treatments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Treatment ID"
// @Param treatment body domain.Treatment true "Treatment details"
// @Success 200 {object} domain.Treatment "Treatment updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Treatment not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /treatments/{id} [put]
func (ctrl *TreatmentController) UpdateTreatment(c *gin.Context) {
	treatmentIDStr := c.Param("id")
	treatmentID, err := strconv.ParseUint(treatmentIDStr, 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid treatment ID")
		return
	}

	// Check if treatment exists
	if _, err := ctrl.service.GetTreatmentByID(uint(treatmentID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Treatment not found")
		return
	}

	var treatment domain.Treatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	treatment.TreatmentID = uint(treatmentID)
	if err := ctrl.service.UpdateTreatment(&treatment); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, treatment)
}

// DeleteTreatment deletes a treatment by ID
// @Summary Delete treatment by ID
// @Description Remove a treatment from the system
// @Tags Treatments
// @Security BearerAuth
// @Produce json
// @Param id path int true "Treatment ID"
// @Success 200 {object} map[string]string "Treatment deleted successfully"
// @Failure 400 {object} map[string]string "Invalid treatment ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Treatment not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /treatments/{id} [delete]
func (ctrl *TreatmentController) DeleteTreatment(c *gin.Context) {
	treatmentIDStr := c.Param("id")
	treatmentID, err := strconv.ParseUint(treatmentIDStr, 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid treatment ID")
		return
	}

	// Check if treatment exists before attempting to delete
	if _, err := ctrl.service.GetTreatmentByID(uint(treatmentID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Treatment not found")
		return
	}

	if err := ctrl.service.DeleteTreatment(uint(treatmentID)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, "Treatment deleted successfully")
}
