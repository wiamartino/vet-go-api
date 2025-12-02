package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
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
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, appointments)
}

func (ctrl *AppointmentController) CreateAppointment(c *gin.Context) {
	var appointment domain.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.service.CreateAppointment(&appointment); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, appointment)
}

func (ctrl *AppointmentController) FindAppointment(c *gin.Context) {

	appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}
	appointment, err := ctrl.service.GetAppointmentByID(uint(appointmentID))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, appointment)
}

func (ctrl *AppointmentController) UpdateAppointment(c *gin.Context) {
	var appointment domain.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}

	// Check if appointment exists
	if _, err := ctrl.service.GetAppointmentByID(uint(appointmentID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Appointment not found")
		return
	}

	appointment.AppointmentID = uint(appointmentID)

	if err := ctrl.service.UpdateAppointment(&appointment); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, appointment)
}

func (ctrl *AppointmentController) DeleteAppointment(c *gin.Context) {

	appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid appointment ID")
		return
	}
	if err := ctrl.service.DeleteAppointment(uint(appointmentID)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Appointment deleted")

}
