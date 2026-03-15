package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicationController struct {
	service *application.MedicationService
}

func NewMedicationController(service *application.MedicationService) *MedicationController {
	return &MedicationController{service: service}
}

// FindMedications retrieves all medications
// @Summary Get all medications
// @Description Get a list of all available medications
// @Tags Medications
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Medication "List of medications"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /medications [get]
func (ctrl *MedicationController) FindMedications(c *gin.Context) {
	medications, err := ctrl.service.GetAllMedications()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, medications)
}

// CreateMedication creates a new medication
// @Summary Create a new medication
// @Description Add a new medication to the system
// @Tags Medications
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param medication body domain.Medication true "Medication details"
// @Success 200 {object} domain.Medication "Medication created successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /medications [post]
func (ctrl *MedicationController) CreateMedication(c *gin.Context) {

	var medication domain.Medication
	if err := c.ShouldBindJSON(&medication); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := ctrl.service.CreateMedication(&medication)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, medication)
}

// FindMedication retrieves a medication by ID
// @Summary Get medication by ID
// @Description Get detailed information about a specific medication
// @Tags Medications
// @Security BearerAuth
// @Produce json
// @Param id path int true "Medication ID"
// @Success 200 {object} domain.Medication "Medication details"
// @Failure 400 {object} map[string]string "Invalid medication ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Medication not found"
// @Router /medications/{id} [get]
func (ctrl *MedicationController) FindMedication(c *gin.Context) {

	medicationID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid medication ID")
		return
	}

	medication, err := ctrl.service.GetMedicationByID(uint(medicationID))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Medication not found")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, medication)

}

// UpdateMedication updates a medication by ID
// @Summary Update medication by ID
// @Description Update information for a specific medication
// @Tags Medications
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Medication ID"
// @Param medication body domain.Medication true "Medication details"
// @Success 200 {object} domain.Medication "Medication updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Medication not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /medications/{id} [put]
func (ctrl *MedicationController) UpdateMedication(c *gin.Context) {

	var medication domain.Medication

	if err := c.ShouldBindJSON(&medication); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	medicationID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid medication ID")
		return
	}

	// Check if medication exists
	if _, err := ctrl.service.GetMedicationByID(uint(medicationID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Medication not found")
		return
	}

	medication.ID = uint(medicationID)
	if err := ctrl.service.UpdateMedication(&medication); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, medication)

}

// DeleteMedication deletes a medication by ID
// @Summary Delete medication by ID
// @Description Remove a medication from the system
// @Tags Medications
// @Security BearerAuth
// @Produce json
// @Param id path int true "Medication ID"
// @Success 200 {object} map[string]string "Medication deleted successfully"
// @Failure 400 {object} map[string]string "Invalid medication ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /medications/{id} [delete]
func (ctrl *MedicationController) DeleteMedication(c *gin.Context) {

	medicationID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid medication ID")
		return
	}

	if err = ctrl.service.DeleteMedication(uint(medicationID)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Medication deleted successfully")
}
