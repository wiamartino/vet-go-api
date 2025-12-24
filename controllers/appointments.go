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

// FindAppointments retrieves all appointments
// @Summary Get all appointments
// @Description Get a list of all appointments
// @Tags Appointments
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Appointment "List of appointments"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /appointments [get]
func (ctrl *AppointmentController) FindAppointments(c *gin.Context) {
	appointments, err := ctrl.service.GetAllAppointments()
	if err != nil {
		utils.RespondWithAppError(c, utils.AsAppError(err))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, appointments)
}

// CreateAppointment creates a new appointment
// @Summary Create a new appointment
// @Description Schedule a new appointment for a pet with a veterinarian
// @Tags Appointments
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param appointment body domain.Appointment true "Appointment details"
// @Success 201 {object} domain.Appointment "Appointment created successfully"
// @Failure 400 {object} map[string]interface{} "Invalid request format or validation errors"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /appointments [post]
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

// FindAppointment retrieves an appointment by ID
// @Summary Get appointment by ID
// @Description Get detailed information about a specific appointment
// @Tags Appointments
// @Security BearerAuth
// @Produce json
// @Param id path int true "Appointment ID"
// @Success 200 {object} domain.Appointment "Appointment details"
// @Failure 400 {object} map[string]string "Invalid appointment ID format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Appointment not found"
// @Router /appointments/{id} [get]
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
