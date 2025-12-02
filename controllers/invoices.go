package controllers

import (
	"go-vet/application"
	"go-vet/domain"
	"go-vet/utils"
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
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithData(c, http.StatusOK, invoices)
}

func (ctrl *InvoiceController) CreateInvoice(c *gin.Context) {
	var invoice domain.Invoice
	if err := c.ShouldBindJSON(&invoice); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	err := ctrl.service.CreateInvoice(&invoice)
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithData(c, http.StatusOK, invoice)
}

func (ctrl *InvoiceController) FindInvoice(c *gin.Context) {

	invoiceID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	invoice, err := ctrl.service.GetInvoiceByID(uint(invoiceID))
	if err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Invoice not found")
		return
	}

	utils.RespondWithData(c, http.StatusOK, invoice)
}

func (ctrl *InvoiceController) UpdateInvoice(c *gin.Context) {

	var invoice domain.Invoice

	if err := c.ShouldBindJSON(&invoice); err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, err.Error())
		return
	}

	invoiceID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	// Check if invoice exists
	if _, err := ctrl.service.GetInvoiceByID(uint(invoiceID)); err != nil {
		utils.RespondWithError(c, http.StatusNotFound, "Invoice not found")
		return
	}

	invoice.InvoiceID = uint(invoiceID)
	if err := ctrl.service.UpdateInvoice(&invoice); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithData(c, http.StatusOK, invoice)
}

func (ctrl *InvoiceController) DeleteInvoice(c *gin.Context) {

	invoiceID, err := strconv.ParseUint(c.Param("id"), 10, 32)
	if err != nil {
		utils.RespondWithError(c, http.StatusBadRequest, "Invalid invoice ID")
		return
	}

	if err = ctrl.service.DeleteInvoice(uint(invoiceID)); err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, "Invoice deleted successfully")

}
