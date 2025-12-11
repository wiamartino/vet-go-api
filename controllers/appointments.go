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
		utils.RespondWithAppError(c, utils.AsAppError(err))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, appointments)
}

func (ctrl *AppointmentController) CreateAppointment(c *gin.Context) {
	var appointment domain.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	// Validate required fields
	validator := utils.NewValidator()
	validator.ValidateNumericID("pet_id", appointment.PetID)
	validator.ValidateNumericID("veterinarian_id", appointment.VeterinarianID)
	validator.ValidateRequired("reason_for_appointment", appointment.ReasonForAppointment)

	if !validator.IsValid() {
		utils.RespondWithValidationError(c, validator.Errors)
		return
	}

	if err := ctrl.service.CreateAppointment(&appointment); err != nil {
		if utils.IsAppError(err) {
			utils.RespondWithAppError(c, err.(*utils.AppError))
		} else {
			utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		}
		return
	}

	utils.RespondWithCreated(c, appointment)
}

func (ctrl *AppointmentController) FindAppointment(c *gin.Context) {
	appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid appointment ID format"))
		return
	}
	appointment, err := ctrl.service.GetAppointmentByID(uint(appointmentID))
	if err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Appointment"))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, appointment)
}

func (ctrl *AppointmentController) UpdateAppointment(c *gin.Context) {
	var appointment domain.Appointment

	if err := c.ShouldBindJSON(&appointment); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid appointment ID format"))
		return
	}

	// Check if appointment exists
	if _, err := ctrl.service.GetAppointmentByID(uint(appointmentID)); err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Appointment"))
		return
	}

	// Validate required fields
	validator := utils.NewValidator()
	validator.ValidateNumericID("pet_id", appointment.PetID)
	validator.ValidateNumericID("veterinarian_id", appointment.VeterinarianID)

	if !validator.IsValid() {
		utils.RespondWithValidationError(c, validator.Errors)
		return
	}

	appointment.AppointmentID = uint(appointmentID)

	if err := ctrl.service.UpdateAppointment(&appointment); err != nil {
		if utils.IsAppError(err) {
			utils.RespondWithAppError(c, err.(*utils.AppError))
		} else {
			utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		}
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, appointment)
}

func (ctrl *AppointmentController) DeleteAppointment(c *gin.Context) {
	appointmentID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid appointment ID format"))
		return
	}

	// Check if appointment exists before attempting to delete
	if _, err := ctrl.service.GetAppointmentByID(uint(appointmentID)); err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Appointment"))
		return
	}

	if err := ctrl.service.DeleteAppointment(uint(appointmentID)); err != nil {
		utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		return
	}

	utils.RespondWithNoContent(c)
}
