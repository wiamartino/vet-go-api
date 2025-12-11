package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ClientController struct {
	service *application.ClientService
}

func NewClientController(service *application.ClientService) *ClientController {
	return &ClientController{service: service}
}

func (ctrl *ClientController) FindClients(c *gin.Context) {
	clients, err := ctrl.service.GetAllClients()
	if err != nil {
		utils.RespondWithAppError(c, utils.AsAppError(err))
		return
	}
	utils.RespondWithSuccess(c, http.StatusOK, clients)
}

func (ctrl *ClientController) CreateClient(c *gin.Context) {
	var client domain.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	// Validate client data
	validator := utils.NewValidator()
	validator.ValidateLengthRange("first_name", client.FirstName, 2, 100)
	validator.ValidateLengthRange("last_name", client.LastName, 2, 100)
	validator.ValidateLengthRange("address", client.Address, 5, 255)
	validator.ValidatePhone("phone", client.Phone)
	validator.ValidateEmail("email", client.Email)

	if !validator.IsValid() {
		utils.RespondWithValidationError(c, validator.Errors)
		return
	}

	if err := ctrl.service.CreateClient(&client); err != nil {
		if utils.IsAppError(err) {
			utils.RespondWithAppError(c, err.(*utils.AppError))
		} else {
			utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		}
		return
	}

	utils.RespondWithCreated(c, client)
}

func (ctrl *ClientController) FindClient(c *gin.Context) {
	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid client ID format"))
		return
	}

	client, err := ctrl.service.GetClientByID(uint(clientID))
	if err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Client"))
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, client)
}

func (ctrl *ClientController) UpdateClient(c *gin.Context) {
	var client domain.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid request format: "+err.Error())
		return
	}

	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid client ID format"))
		return
	}

	// Check if client exists
	if _, err := ctrl.service.GetClientByID(uint(clientID)); err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Client"))
		return
	}

	// Validate client data
	validator := utils.NewValidator()
	validator.ValidateLengthRange("first_name", client.FirstName, 2, 100)
	validator.ValidateLengthRange("last_name", client.LastName, 2, 100)
	validator.ValidateLengthRange("address", client.Address, 5, 255)
	validator.ValidatePhone("phone", client.Phone)
	validator.ValidateEmail("email", client.Email)

	if !validator.IsValid() {
		utils.RespondWithValidationError(c, validator.Errors)
		return
	}

	client.ClientID = uint(clientID)

	if err := ctrl.service.UpdateClient(&client); err != nil {
		if utils.IsAppError(err) {
			utils.RespondWithAppError(c, err.(*utils.AppError))
		} else {
			utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		}
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, client)
}

func (ctrl *ClientController) DeleteClient(c *gin.Context) {
	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithAppError(c, utils.NewBadRequestError("Invalid client ID format"))
		return
	}

	// Check if client exists before attempting to delete
	if _, err := ctrl.service.GetClientByID(uint(clientID)); err != nil {
		utils.RespondWithAppError(c, utils.NewNotFoundError("Client"))
		return
	}

	if err = ctrl.service.DeleteClient(uint(clientID)); err != nil {
		utils.RespondWithAppError(c, utils.NewInternalError(err.Error()))
		return
	}

	utils.RespondWithNoContent(c)
}
