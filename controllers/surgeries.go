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

func (c *SurgeryController) GetSurgeries(ctx *gin.Context) {
	surgeries, err := c.service.GetAllSurgeries()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

func (c *SurgeryController) GetSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	surgery, err := c.service.GetSurgeryByID(uint(id))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgery)
}

func (c *SurgeryController) GetSurgeriesByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	surgeries, err := c.service.GetSurgeriesByPetID(uint(petID))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

func (c *SurgeryController) GetSurgeriesByStatus(ctx *gin.Context) {
	statusStr := ctx.Query("status")
	if statusStr == "" {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Status parameter is required")
		return
	}

	status := domain.SurgeryStatus(statusStr)
	surgeries, err := c.service.GetSurgeriesByStatus(status)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

func (c *SurgeryController) GetScheduledSurgeries(ctx *gin.Context) {
	startDateStr := ctx.Query("start_date")
	endDateStr := ctx.Query("end_date")

	if startDateStr == "" || endDateStr == "" {
		utils.RespondWithError(ctx, http.StatusBadRequest, "start_date and end_date are required")
		return
	}

	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid start_date format. Use YYYY-MM-DD")
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, "Invalid end_date format. Use YYYY-MM-DD")
		return
	}

	surgeries, err := c.service.GetScheduledSurgeries(startDate, endDate)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgeries)
}

func (c *SurgeryController) CreateSurgery(ctx *gin.Context) {
	var surgery domain.Surgery
	if err := ctx.ShouldBindJSON(&surgery); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CreateSurgery(&surgery); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusCreated, surgery)
}

func (c *SurgeryController) UpdateSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var surgery domain.Surgery
	if err := ctx.ShouldBindJSON(&surgery); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	surgery.SurgeryID = uint(id)
	if err := c.service.UpdateSurgery(&surgery); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, surgery)
}

func (c *SurgeryController) StartSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.StartSurgery(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Surgery started successfully"})
}

func (c *SurgeryController) CompleteSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var body struct {
		PostOpNotes   string `json:"post_op_notes"`
		Complications string `json:"complications"`
	}

	if err := ctx.ShouldBindJSON(&body); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CompleteSurgery(uint(id), body.PostOpNotes, body.Complications); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Surgery completed successfully"})
}

func (c *SurgeryController) CancelSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CancelSurgery(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Surgery cancelled successfully"})
}

func (c *SurgeryController) DeleteSurgery(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.DeleteSurgery(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Surgery deleted successfully"})
}
