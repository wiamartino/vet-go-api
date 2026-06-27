package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type SurgeryController struct {
	service *application.SurgeryService
}

func NewSurgeryController(service *application.SurgeryService) *SurgeryController {
	return &SurgeryController{service: service}
}

// GetSurgeries retrieves all surgeries
// @Summary Get all surgeries
// @Description Get a list of all surgeries
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Surgery "List of surgeries"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries [get]
func (c *SurgeryController) GetSurgeries(ctx *gin.Context) {
	surgeries, err := c.service.GetAllSurgeries()
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}
	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

// GetSurgery retrieves a surgery by ID
// @Summary Get surgery by ID
// @Description Get detailed information about a specific surgery
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Param id path int true "Surgery ID"
// @Success 200 {object} domain.Surgery "Surgery details"
// @Failure 400 {object} map[string]string "Invalid surgery ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Surgery not found"
// @Router /surgeries/{id} [get]
func (c *SurgeryController) GetSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid surgery ID"))
		return
	}

	surgery, err := c.service.GetSurgeryByID(uint(id))
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Surgery"))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgery)
}

// GetSurgeriesByPet retrieves surgeries for a specific pet
// @Summary Get surgeries by pet ID
// @Description Get all surgeries for a specific pet
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Param id path int true "Pet ID"
// @Success 200 {array} domain.Surgery "List of surgeries for the pet"
// @Failure 400 {object} map[string]string "Invalid pet ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /pets/{id}/surgeries [get]
func (c *SurgeryController) GetSurgeriesByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid pet ID"))
		return
	}

	surgeries, err := c.service.GetSurgeriesByPetID(uint(petID))
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

// GetSurgeriesByStatus retrieves surgeries by status
// @Summary Get surgeries by status
// @Description Get all surgeries filtered by status
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Param status query string true "Surgery status (scheduled, in_progress, completed, cancelled)"
// @Success 200 {array} domain.Surgery "List of surgeries with the specified status"
// @Failure 400 {object} map[string]string "Status parameter is required"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries/status [get]
func (c *SurgeryController) GetSurgeriesByStatus(ctx *gin.Context) {
	statusStr := ctx.Query("status")
	if statusStr == "" {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Status parameter is required"))
		return
	}

	status := domain.SurgeryStatus(statusStr)
	surgeries, err := c.service.GetSurgeriesByStatus(status)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

// GetScheduledSurgeries retrieves scheduled surgeries within a date range
// @Summary Get scheduled surgeries
// @Description Get all scheduled surgeries within a specified date range
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Param start_date query string true "Start date (YYYY-MM-DD)"
// @Param end_date query string true "End date (YYYY-MM-DD)"
// @Success 200 {array} domain.Surgery "List of scheduled surgeries"
// @Failure 400 {object} map[string]string "Invalid date format or missing parameters"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries/scheduled [get]
func (c *SurgeryController) GetScheduledSurgeries(ctx *gin.Context) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("start_date and end_date are required"))
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid start_date format. Use YYYY-MM-DD"))
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid end_date format. Use YYYY-MM-DD"))
		return
	}

	surgeries, err := c.service.GetScheduledSurgeries(startDate, endDate)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

// CreateSurgery creates a new surgery
// @Summary Create a new surgery
// @Description Schedule a new surgery for a pet
// @Tags Surgeries
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param surgery body domain.Surgery true "Surgery details"
// @Success 201 {object} domain.Surgery "Surgery created successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries [post]
func (c *SurgeryController) CreateSurgery(ctx *gin.Context) {
	var surgery domain.Surgery
	if err := ctx.ShouldBindJSON(&surgery); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	validator := utils.NewValidator()
	validator.ValidateNumericID("pet_id", surgery.PetID)
	validator.ValidateNumericID("veterinarian_id", surgery.VeterinarianID)
	validator.ValidateRequired("surgery_name", surgery.SurgeryName)
	validator.ValidateInSlice("surgery_type", string(surgery.SurgeryType), []string{"routine", "emergency", "elective"})
	validator.ValidateInSlice("status", string(surgery.Status), []string{"scheduled", "in_progress", "completed", "cancelled"})

	if !validator.IsValid() {
		utils.RespondWithValidationError(ctx, validator.Errors)
		return
	}

	if err := c.service.CreateSurgery(&surgery); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithCreated(ctx, surgery)
}

// UpdateSurgery updates a surgery by ID
// @Summary Update surgery by ID
// @Description Update information for a specific surgery
// @Tags Surgeries
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Surgery ID"
// @Param surgery body domain.Surgery true "Surgery details"
// @Success 200 {object} domain.Surgery "Surgery updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries/{id} [put]
func (c *SurgeryController) UpdateSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid surgery ID"))
		return
	}

	// Check if surgery exists
	if _, err := c.service.GetSurgeryByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Surgery"))
		return
	}

	var surgery domain.Surgery
	if err := ctx.ShouldBindJSON(&surgery); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	validator := utils.NewValidator()
	validator.ValidateNumericID("pet_id", surgery.PetID)
	validator.ValidateNumericID("veterinarian_id", surgery.VeterinarianID)
	validator.ValidateRequired("surgery_name", surgery.SurgeryName)
	validator.ValidateInSlice("surgery_type", string(surgery.SurgeryType), []string{"routine", "emergency", "elective"})
	validator.ValidateInSlice("status", string(surgery.Status), []string{"scheduled", "in_progress", "completed", "cancelled"})

	if !validator.IsValid() {
		utils.RespondWithValidationError(ctx, validator.Errors)
		return
	}

	surgery.SurgeryID = uint(id)
	if err := c.service.UpdateSurgery(&surgery); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgery)
}

// StartSurgery starts a surgery
// @Summary Start a surgery
// @Description Mark a surgery as in progress
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Param id path int true "Surgery ID"
// @Success 200 {object} map[string]string "Surgery started successfully"
// @Failure 400 {object} map[string]string "Invalid surgery ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries/{id}/start [patch]
func (c *SurgeryController) StartSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid surgery ID"))
		return
	}

	if _, err := c.service.GetSurgeryByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Surgery"))
		return
	}

	if err := c.service.StartSurgery(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Surgery started successfully"})
}

// CompleteSurgery completes a surgery
// @Summary Complete a surgery
// @Description Mark a surgery as completed with post-operative notes
// @Tags Surgeries
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Surgery ID"
// @Param body body object{post_op_notes=string,complications=string} true "Completion details"
// @Success 200 {object} map[string]string "Surgery completed successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries/{id}/complete [patch]
func (c *SurgeryController) CompleteSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid surgery ID"))
		return
	}

	if _, err := c.service.GetSurgeryByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Surgery"))
		return
	}

	var body struct {
		PostOpNotes   string `json:"post_op_notes"`
		Complications string `json:"complications"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	if err := c.service.CompleteSurgery(uint(id), body.PostOpNotes, body.Complications); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Surgery completed successfully"})
}

// CancelSurgery cancels a surgery
// @Summary Cancel a surgery
// @Description Mark a surgery as cancelled
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Param id path int true "Surgery ID"
// @Success 200 {object} map[string]string "Surgery cancelled successfully"
// @Failure 400 {object} map[string]string "Invalid surgery ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries/{id}/cancel [patch]
func (c *SurgeryController) CancelSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid surgery ID"))
		return
	}

	if _, err := c.service.GetSurgeryByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Surgery"))
		return
	}

	if err := c.service.CancelSurgery(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Surgery cancelled successfully"})
}

// DeleteSurgery deletes a surgery by ID
// @Summary Delete surgery by ID
// @Description Remove a surgery from the system
// @Tags Surgeries
// @Security BearerAuth
// @Produce json
// @Param id path int true "Surgery ID"
// @Success 204 "Surgery deleted successfully"
// @Failure 400 {object} map[string]string "Invalid surgery ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /surgeries/{id} [delete]
func (c *SurgeryController) DeleteSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(ctx, utils.NewBadRequestError("Invalid surgery ID"))
		return
	}

	if _, err := c.service.GetSurgeryByID(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.NewNotFoundError("Surgery"))
		return
	}

	if err := c.service.DeleteSurgery(uint(id)); err != nil {
		utils.RespondWithAppError(ctx, utils.AsAppError(err))
		return
	}

	utils.RespondWithNoContent(ctx)
}
