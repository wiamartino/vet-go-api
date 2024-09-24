package controllers

import (
	"go-vet/config"
	"go-vet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindMedications(c *gin.Context) {
	var medications []models.Medication
	config.DB.Find(&medications)
	c.JSON(http.StatusOK, medications)
}

func CreateMedication(c *gin.Context) {
	var medication models.Medication
	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := config.DB.Create(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medication)
}

func FindMedication(c *gin.Context) {
	var medication models.Medication
	if err := config.DB.First(&medication, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}
	c.JSON(http.StatusOK, medication)
}

func UpdateMedication(c *gin.Context) {
	var medication models.Medication
	if err := config.DB.First(&medication, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}
	if err := c.ShouldBindJSON(&medication); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, medication)
}

func DeleteMedication(c *gin.Context) {
	var medication models.Medication
	if err := config.DB.First(&medication, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Medication not found"})
		return
	}
	if err := config.DB.Delete(&medication).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Medication deleted"})
}
