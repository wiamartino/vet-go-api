package controllers

import (
	"go-vet/application"
	"go-vet/domain"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch veterinarians"})
		return
	}
	c.JSON(http.StatusOK, veterinarians)
}

func (ctrl *VeterinarianController) CreateVeterinarian(c *gin.Context) {
	var veterinarian domain.Veterinarian
	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.CreateVeterinarian(&veterinarian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create veterinarian"})
		return
	}
	c.JSON(http.StatusOK, veterinarian)
}

func (ctrl *VeterinarianController) FindVeterinarian(c *gin.Context) {
	idParam, _ := strconv.Atoi(c.Param("id"))
	id := uint(idParam)
	veterinarian, err := ctrl.service.GetVeterinarianByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, veterinarian)
}

func (ctrl *VeterinarianController) UpdateVeterinarian(c *gin.Context) {

	var veterinarian domain.Veterinarian

	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	veterinarianID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	veterinarian.VeterinarianID = uint(veterinarianID)
	if err := ctrl.service.UpdateVeterinarian(&veterinarian); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, veterinarian)
}

func (ctrl *VeterinarianController) DeleteVeterinarian(c *gin.Context) {

	idParam, _ := strconv.Atoi(c.Param("id"))
	id := uint(idParam)
	if _, err := ctrl.service.GetVeterinarianByID(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Veterinarian not found"})
		return
	}

	if err := ctrl.service.DeleteVeterinarian(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete veterinarian"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Veterinarian deleted"})
}
