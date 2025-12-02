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

func (ctrl *VeterinarianController) FindVeterinarians(c *gin.Context) {
	veterinarians, err := ctrl.service.GetAllVeterinarians()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to fetch veterinarians")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, veterinarians)
}

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
	veterinarian.VeterinarianID = uint(veterinarianID)
	if err := ctrl.service.UpdateVeterinarian(&veterinarian); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, "Failed to update veterinarian")
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, veterinarian)
}

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
