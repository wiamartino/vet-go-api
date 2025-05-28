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

func (ctrl *MedicationController) FindMedications(c *gin.Context) {
	medications, err := ctrl.service.GetAllMedications()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, medications)
}

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
	medication.ID = uint(medicationID)
	if err := ctrl.service.UpdateMedication(&medication); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, medication)

}

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
