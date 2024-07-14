package controllers

import (
	"go-vet/config"
	"go-vet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindTreatments(c *gin.Context) {
	var treatments []models.Treatment
	config.DB.Find(&treatments)
	c.JSON(http.StatusOK, treatments)
}

func CreateTreatment(c *gin.Context) {
	var treatment models.Treatment
	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&treatment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create treatment"})
		return
	}
	c.JSON(http.StatusOK, treatment)
}

func FindTreatment(c *gin.Context) {
	var treatment models.Treatment
	if err := config.DB.First(&treatment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Treatment not found"})
		return
	}
	c.JSON(http.StatusOK, treatment)
}

func UpdateTreatment(c *gin.Context) {
	var treatment models.Treatment
	if err := config.DB.First(&treatment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Treatment not found"})
		return
	}
	if err := c.ShouldBindJSON(&treatment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&treatment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update treatment"})
		return
	}
	c.JSON(http.StatusOK, treatment)
}

func DeleteTreatment(c *gin.Context) {
	var treatment models.Treatment
	if err := config.DB.First(&treatment, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Treatment not found"})
		return
	}
	if err := config.DB.Delete(&treatment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete treatment"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Treatment deleted"})
}
