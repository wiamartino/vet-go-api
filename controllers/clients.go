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
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithData(c, http.StatusOK, clients)
}

func (ctrl *ClientController) CreateClient(c *gin.Context) {
	var client domain.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := ctrl.service.CreateClient(&client); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithData(c, http.StatusOK, client)
}

func (ctrl *ClientController) FindClient(c *gin.Context) {

	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid client ID")
		return
	}

	client, err := ctrl.service.GetClientByID(uint(clientID))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Client not found")
		return
	}

	utils.RespondWithData(c, http.StatusOK, client)
}

func (ctrl *ClientController) UpdateClient(c *gin.Context) {

	var client domain.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid client ID")
		return
	}
	client.ClientID = uint(clientID)

	if err := ctrl.service.UpdateClient(&client); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithData(c, http.StatusOK, client)

}

func (ctrl *ClientController) DeleteClient(c *gin.Context) {

	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid client ID")
		return
	}

	if err = ctrl.service.DeleteClient(uint(clientID)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Client deleted successfully")
}
