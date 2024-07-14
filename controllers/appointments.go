package controllers

import (
	"go-vet/config"
	"go-vet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindAppointments(c *gin.Context) {
	var appointments []models.Appointment
	config.DB.Preload("Pet").Preload("Veterinarian").Find(&appointments)
	c.JSON(http.StatusOK, appointments)
}

func CreateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func FindAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := config.DB.Preload("Pet").Preload("Veterinarian").First(&appointment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}
	c.JSON(http.StatusOK, appointment)
}

func UpdateAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := config.DB.First(&appointment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}
	if err := c.ShouldBindJSON(&appointment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Save(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, appointment)
}

func DeleteAppointment(c *gin.Context) {
	var appointment models.Appointment
	if err := config.DB.First(&appointment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Appointment not found"})
		return
	}
	if err := config.DB.Delete(&appointment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Appointment deleted"})
}
