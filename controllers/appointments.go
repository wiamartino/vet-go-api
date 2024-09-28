package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AppointmentController struct {
	service *application.AppointmentService
}

func NewAppointmentController(service *application.AppointmentService) *AppointmentController {
	return &AppointmentController{service: service}
}

func (ctrl *AppointmentController) FindAppointments(c *gin.Context) {
	appointments, err := ctrl.service.GetAllAppointments()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointments)
}

func (ctrl *AppointmentController) CreateAppointment(c *gin.Context) {
	var appointment domain.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateAppointment(&appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ctrl *AppointmentController) FindAppointment(c *gin.Context) {

	idParam, _ := strconv.Atoi(c.Param("id"))
	id := uint(idParam)
	appointment, err := ctrl.service.GetAppointmentByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, appointment)
}

func (ctrl *AppointmentController) UpdateAppointment(c *gin.Context) {
	var appointment domain.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.UpdateAppointment(&appointment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func (ctrl *AppointmentController) DeleteAppointment(c *gin.Context) {

	idParam, _ := strconv.Atoi(c.Param("id"))
	id := uint(idParam)
	if err := ctrl.service.DeleteAppointment(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Appointment deleted"})

}
