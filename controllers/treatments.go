package controllers

import (
	"go-vet/application"
	"go-vet/domain"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, treatments)
}

func (ctrl *TreatmentController) CreateTreatment(c *gin.Context) {

	var treatment domain.Treatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.CreateTreatment(&treatment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, treatment)
}

func (ctrl *TreatmentController) FindTreatment(c *gin.Context) {

	treatmentID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	treatment, err := ctrl.service.GetTreatmentByID(uint(treatmentID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Treatment not found"})
		return
	}
	c.JSON(http.StatusOK, treatment)
}

func (ctrl *TreatmentController) UpdateTreatment(c *gin.Context) {

	var treatment domain.Treatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.UpdateTreatment(&treatment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, treatment)
}

func (ctrl *TreatmentController) DeleteTreatment(c *gin.Context) {

	treatmentID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := ctrl.service.DeleteTreatment(uint(treatmentID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Treatment deleted"})
}
