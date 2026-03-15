package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type VeterinarianController struct {
	service *application.VeterinarianService
}

func NewVeterinarianController(service *application.VeterinarianService) *VeterinarianController {
	return &VeterinarianController{service: service}
}

// FindVeterinarians retrieves all veterinarians
// @Summary Get all veterinarians
// @Description Get a list of all veterinarians in the clinic
// @Tags Veterinarians
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Veterinarian "List of veterinarians"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /veterinarians [get]
func (ctrl *VeterinarianController) FindVeterinarians(c *gin.Context) {
	veterinarians, err := ctrl.service.GetAllVeterinarians()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch veterinarians")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, veterinarians)
}

// CreateVeterinarian creates a new veterinarian
// @Summary Create a new veterinarian
// @Description Add a new veterinarian to the clinic
// @Tags Veterinarians
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param veterinarian body domain.Veterinarian true "Veterinarian details"
// @Success 200 {object} domain.Veterinarian "Veterinarian created successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /veterinarians [post]
func (ctrl *VeterinarianController) CreateVeterinarian(c *gin.Context) {
	var veterinarian domain.Veterinarian
	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctrl.service.CreateVeterinarian(&veterinarian); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to create veterinarian")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, veterinarian)
}

// FindVeterinarian retrieves a veterinarian by ID
// @Summary Get veterinarian by ID
// @Description Get detailed information about a specific veterinarian
// @Tags Veterinarians
// @Security BearerAuth
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Success 200 {object} domain.Veterinarian "Veterinarian details"
// @Failure 400 {object} map[string]string "Invalid veterinarian ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /veterinarians/{id} [get]
func (ctrl *VeterinarianController) FindVeterinarian(c *gin.Context) {
	veterinarianID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	veterinarian, err := ctrl.service.GetVeterinarianByID(uint(veterinarianID))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, veterinarian)
}

// UpdateVeterinarian updates a veterinarian by ID
// @Summary Update veterinarian by ID
// @Description Update information for a specific veterinarian
// @Tags Veterinarians
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Param veterinarian body domain.Veterinarian true "Veterinarian details"
// @Success 200 {object} domain.Veterinarian "Veterinarian updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Veterinarian not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /veterinarians/{id} [put]
func (ctrl *VeterinarianController) UpdateVeterinarian(c *gin.Context) {

	var veterinarian domain.Veterinarian

	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	veterinarianID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}

	// Check if veterinarian exists
	if _, err := ctrl.service.GetVeterinarianByID(uint(veterinarianID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Veterinarian not found")
		return
	}

	veterinarian.VeterinarianID = uint(veterinarianID)
	if err := ctrl.service.UpdateVeterinarian(&veterinarian); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update veterinarian")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, veterinarian)
}

// DeleteVeterinarian deletes a veterinarian by ID
// @Summary Delete veterinarian by ID
// @Description Remove a veterinarian from the system
// @Tags Veterinarians
// @Security BearerAuth
// @Produce json
// @Param id path int true "Veterinarian ID"
// @Success 200 {object} map[string]string "Veterinarian deleted successfully"
// @Failure 400 {object} map[string]string "Invalid veterinarian ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Veterinarian not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /veterinarians/{id} [delete]
func (ctrl *VeterinarianController) DeleteVeterinarian(c *gin.Context) {

	veterinarianID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid ID")
		return
	}
	id := uint(veterinarianID)
	if _, err := ctrl.service.GetVeterinarianByID(id); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Veterinarian not found")
		return
	}

	if err := ctrl.service.DeleteVeterinarian(id); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to delete veterinarian")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Veterinarian deleted")
}
