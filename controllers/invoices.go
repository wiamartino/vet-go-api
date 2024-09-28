package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InvoiceController struct {
	service *application.InvoiceService
}

func NewInvoiceController(service *application.InvoiceService) *InvoiceController {
	return &InvoiceController{service: service}
}

func (ctrl *InvoiceController) FindInvoices(c *gin.Context) {
	invoices, err := ctrl.service.GetAllInvoices()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoices)
}

func (ctrl *InvoiceController) CreateInvoice(c *gin.Context) {
	var invoice domain.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ctrl.service.CreateInvoice(invoice)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (ctrl *InvoiceController) FindInvoice(c *gin.Context) {

	invoiceID, _ := strconv.ParseUint(c.Param("id"), 10, 32)

	invoice, err := ctrl.service.GetInvoiceByID(uint(invoiceID))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Invoice not found"})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (ctrl *InvoiceController) UpdateInvoice(c *gin.Context) {

	var invoice domain.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := ctrl.service.UpdateInvoice(invoice); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, invoice)
}

func (ctrl *InvoiceController) DeleteInvoice(c *gin.Context) {

	invoiceID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid invoice ID"})
		return
	}

	if err = ctrl.service.DeleteInvoice(uint(invoiceID)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Invoice deleted"})

}
