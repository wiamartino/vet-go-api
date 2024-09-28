package controllers

import (
	"go-vet/application"
	"go-vet/domain"
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, clients)
}

func (ctrl *ClientController) CreateClient(c *gin.Context) {
	var client domain.Client
	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.CreateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (ctrl *ClientController) FindClient(c *gin.Context) {

	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	client, err := ctrl.service.GetClientByID(uint(clientID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Client not found"})
		return
	}

	c.JSON(http.StatusOK, client)
}

func (ctrl *ClientController) UpdateClient(c *gin.Context) {

	var client domain.Client

	if err := c.ShouldBindJSON(&client); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clientID, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	client.ClientID = uint(clientID)

	if err := ctrl.service.UpdateClient(&client); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, client)

}

func (ctrl *ClientController) DeleteClient(c *gin.Context) {

	clientID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid client ID"})
		return
	}

	if err = ctrl.service.DeleteClient(uint(clientID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Client deleted"})
}
