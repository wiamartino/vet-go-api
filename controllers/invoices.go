package controllers

import (
	"go-vet/config"
	"go-vet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindInvoices(c *gin.Context) {
	var invoices []models.Invoice
	config.DB.Preload("Client").Preload("Appointment").Find(&invoices)
	c.JSON(http.StatusOK, invoices)
}

func CreateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create invoice"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func FindInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := config.DB.Preload("Client").Preload("Appointment").First(&invoice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func UpdateInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := config.DB.First(&invoice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Save(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update invoice"})
		return
	}
	c.JSON(http.StatusOK, invoice)
}

func DeleteInvoice(c *gin.Context) {
	var invoice models.Invoice
	if err := config.DB.First(&invoice, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}
	if err := config.DB.Delete(&invoice).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete invoice"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "Invoice deleted"})
}
