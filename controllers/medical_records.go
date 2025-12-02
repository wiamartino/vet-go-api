package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type MedicalRecordController struct {
	service *application.MedicalRecordService
}

func NewMedicalRecordController(service *application.MedicalRecordService) *MedicalRecordController {
	return &MedicalRecordController{service: service}
}

// GetMedicalRecords godoc
// @Summary Get all medical records
// @Description Get all medical records
// @Tags medical-records
// @Accept json
// @Produce json
// @Success 200 {array} domain.MedicalRecord
// @Router /api/v1/medical-records [get]
func (c *MedicalRecordController) GetMedicalRecords(ctx *gin.Context) {
	records, err := c.service.GetAllRecords()
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(ctx, http.StatusOK, records)
}

// GetMedicalRecord godoc
// @Summary Get medical record by ID
// @Description Get a medical record by ID
// @Tags medical-records
// @Accept json
// @Produce json
// @Param id path int true "Medical Record ID"
// @Success 200 {object} domain.MedicalRecord
// @Router /api/v1/medical-records/{id} [get]
func (c *MedicalRecordController) GetMedicalRecord(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	record, err := c.service.GetRecordByID(uint(id))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusNotFound, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, record)
}

// GetMedicalRecordsByPet godoc
// @Summary Get medical records by pet ID
// @Description Get all medical records for a specific pet
// @Tags medical-records
// @Accept json
// @Produce json
// @Param pet_id path int true "Pet ID"
// @Success 200 {array} domain.MedicalRecord
// @Router /api/v1/pets/{pet_id}/medical-records [get]
func (c *MedicalRecordController) GetMedicalRecordsByPet(ctx *gin.Context) {
	petID, err := strconv.ParseUint(ctx.Param("pet_id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	records, err := c.service.GetRecordsByPetID(uint(petID))
	if err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, records)
}

// CreateMedicalRecord godoc
// @Summary Create a new medical record
// @Description Create a new medical record
// @Tags medical-records
// @Accept json
// @Produce json
// @Param record body domain.MedicalRecord true "Medical Record"
// @Success 201 {object} domain.MedicalRecord
// @Router /api/v1/medical-records [post]
func (c *MedicalRecordController) CreateMedicalRecord(ctx *gin.Context) {
	var record domain.MedicalRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.CreateRecord(&record); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusCreated, record)
}

// UpdateMedicalRecord godoc
// @Summary Update a medical record
// @Description Update an existing medical record
// @Tags medical-records
// @Accept json
// @Produce json
// @Param id path int true "Medical Record ID"
// @Param record body domain.MedicalRecord true "Medical Record"
// @Success 200 {object} domain.MedicalRecord
// @Router /api/v1/medical-records/{id} [put]
func (c *MedicalRecordController) UpdateMedicalRecord(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	var record domain.MedicalRecord
	if err := ctx.ShouldBindJSON(&record); err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	record.MedicalRecordID = uint(id)
	if err := c.service.UpdateRecord(&record); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, record)
}

// DeleteMedicalRecord godoc
// @Summary Delete a medical record
// @Description Delete a medical record by ID
// @Tags medical-records
// @Accept json
// @Produce json
// @Param id path int true "Medical Record ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/v1/medical-records/{id} [delete]
func (c *MedicalRecordController) DeleteMedicalRecord(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	if err := c.service.DeleteRecord(uint(id)); err != nil {
		utils.RespondWithError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(ctx, http.StatusOK, gin.H{"message": "Medical record deleted successfully"})
}
