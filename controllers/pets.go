package controllers

import (
	"go-vet/config"
	"go-vet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindPets(c *gin.Context) {
	var pets []models.Pet
	config.DB.Preload("Client").Find(&pets)
	c.JSON(http.StatusOK, pets)
}

func CreatePet(c *gin.Context) {
	var pet models.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	config.DB.Create(&pet)
	c.JSON(http.StatusOK, pet)
}

func FindPet(c *gin.Context) {
	var pet models.Pet
	if err := config.DB.Preload("Client").First(&pet, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func UpdatePet(c *gin.Context) {
	var pet models.Pet
	if err := config.DB.First(&pet, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&pet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pet"})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func DeletePet(c *gin.Context) {
	var pet models.Pet
	if err := config.DB.First(&pet, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}
	if err := config.DB.Delete(&pet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pet"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Pet deleted"})
}
