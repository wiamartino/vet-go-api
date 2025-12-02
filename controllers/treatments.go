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

func (ctrl *TreatmentController) FindTreatments(c *gin.Context) {
	treatments, err := ctrl.service.GetAllTreatments()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, treatments)
}

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
