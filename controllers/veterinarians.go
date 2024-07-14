package controllers

import (
	"go-vet/config"
	"go-vet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindVeterinarians(c *gin.Context) {
	var veterinarians []models.Veterinarian
	config.DB.Find(&veterinarians)
	c.JSON(http.StatusOK, veterinarians)
}

func CreateVeterinarian(c *gin.Context) {
	var veterinarian models.Veterinarian
	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&veterinarian).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create veterinarian"})
		return
	}
	c.JSON(http.StatusOK, veterinarian)
}

func FindVeterinarian(c *gin.Context) {
	var veterinarian models.Veterinarian
	if err := config.DB.First(&veterinarian, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Veterinarian not found"})
		return
	}
	c.JSON(http.StatusOK, veterinarian)
}

func UpdateVeterinarian(c *gin.Context) {
	var veterinarian models.Veterinarian
	if err := config.DB.First(&veterinarian, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Veterinarian not found"})
		return
	}
	if err := c.ShouldBindJSON(&veterinarian); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&veterinarian).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update veterinarian"})
		return
	}
	c.JSON(http.StatusOK, veterinarian)
}

func DeleteVeterinarian(c *gin.Context) {
	var veterinarian models.Veterinarian
	if err := config.DB.First(&veterinarian, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Veterinarian not found"})
		return
	}
	if err := config.DB.Delete(&veterinarian).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete veterinarian"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Veterinarian deleted"})
}
