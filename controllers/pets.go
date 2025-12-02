package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
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
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, pets)
}

func (ctrl *PetController) CreatePet(c *gin.Context) {

	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := ctrl.service.CreatePet(&pet); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, pet)
}

func (ctrl *PetController) FindPet(c *gin.Context) {

	petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid pet ID")
		return
	}
	pet, err := ctrl.service.GetPetByID(uint(petID))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Pet not found")
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, pet)
}

func (ctrl *PetController) UpdatePet(c *gin.Context) {

	var pet domain.Pet
	if err := c.ShouldBindJSON(&pet); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid pet ID")
		return
	}

	// Check if pet exists
	if _, err := ctrl.service.GetPetByID(uint(petID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Pet not found")
		return
	}

	pet.PetID = uint(petID)
	err = ctrl.service.UpdatePet(&pet)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, pet)
}

func (ctrl *PetController) DeletePet(c *gin.Context) {

	petID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid pet ID")
		return
	}

	// Check if pet exists before attempting to delete
	if _, err := ctrl.service.GetPetByID(uint(petID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Pet not found")
		return
	}

	err = ctrl.service.DeletePet(uint(petID))
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Pet deleted")
}
