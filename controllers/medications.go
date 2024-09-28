package controllers

import (
	"go-vet/application"
	"go-vet/domain"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, medications)
}

func (ctrl *MedicationController) CreateMedication(c *gin.Context) {

	var medication domain.Medication
	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.CreateMedication(&medication)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medication)
}

func (ctrl *MedicationController) FindMedication(c *gin.Context) {

	medicationID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	medication, err := ctrl.service.GetMedicationByID(uint(medicationID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}

	c.JSON(http.StatusOK, medication)

}

func (ctrl *MedicationController) UpdateMedication(c *gin.Context) {

	var medication domain.Medication

	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.UpdateMedication(&medication); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medication)

}

func (ctrl *MedicationController) DeleteMedication(c *gin.Context) {

	medicationID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid medication ID"})
		return
	}

	if err = ctrl.service.DeleteMedication(uint(medicationID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Medication deleted"})
}
