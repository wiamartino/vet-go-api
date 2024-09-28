package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PetController struct {
	service *application.PetService
}

func NewPetController(service *application.PetService) *PetController {
	return &PetController{service: service}
}

func (ctrl *PetController) FindPets(c *gin.Context) {
	pets, err := ctrl.service.GetAllPets()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pets)
}

func (ctrl *PetController) CreatePet(c *gin.Context) {

	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := ctrl.service.CreatePet(&pet); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func (ctrl *PetController) FindPet(c *gin.Context) {

	petID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	pet, err := ctrl.service.GetPetByID(uint(petID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		return
	}
	c.JSON(http.StatusOK, pet)
}

func (ctrl *PetController) UpdatePet(c *gin.Context) {

	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.UpdatePet(&pet)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, pet)
}

func (ctrl *PetController) DeletePet(c *gin.Context) {

	petID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	err := ctrl.service.DeletePet(uint(petID))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Pet deleted"})
}
