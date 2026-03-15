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

// FindInvoices retrieves all invoices
// @Summary Get all invoices
// @Description Get a list of all invoices
// @Tags Invoices
// @Security BearerAuth
// @Produce json
// @Success 200 {array} domain.Invoice "List of invoices"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoices [get]
func (ctrl *InvoiceController) FindInvoices(c *gin.Context) {
	invoices, err := ctrl.service.GetAllInvoices()
	if err != nil {
		utils.RespondWithError(c, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithSuccess(c, http.StatusOK, invoices)
}

// CreateInvoice creates a new invoice
// @Summary Create a new invoice
// @Description Create a new invoice for a client appointment
// @Tags Invoices
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param invoice body domain.Invoice true "Invoice details"
// @Success 200 {object} domain.Invoice "Invoice created successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoices [post]
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

	utils.RespondWithSuccess(c, http.StatusOK, invoice)
}

// FindInvoice retrieves an invoice by ID
// @Summary Get invoice by ID
// @Description Get detailed information about a specific invoice
// @Tags Invoices
// @Security BearerAuth
// @Produce json
// @Param id path int true "Invoice ID"
// @Success 200 {object} domain.Invoice "Invoice details"
// @Failure 400 {object} map[string]string "Invalid invoice ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Invoice not found"
// @Router /invoices/{id} [get]
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

	utils.RespondWithSuccess(c, http.StatusOK, invoice)
}

// UpdateInvoice updates an invoice by ID
// @Summary Update invoice by ID
// @Description Update information for a specific invoice
// @Tags Invoices
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param id path int true "Invoice ID"
// @Param invoice body domain.Invoice true "Invoice details"
// @Success 200 {object} domain.Invoice "Invoice updated successfully"
// @Failure 400 {object} map[string]string "Invalid request format"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 404 {object} map[string]string "Invoice not found"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoices/{id} [put]
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

	utils.RespondWithSuccess(c, http.StatusOK, invoice)
}

// DeleteInvoice deletes an invoice by ID
// @Summary Delete invoice by ID
// @Description Remove an invoice from the system
// @Tags Invoices
// @Security BearerAuth
// @Produce json
// @Param id path int true "Invoice ID"
// @Success 200 {object} map[string]string "Invoice deleted successfully"
// @Failure 400 {object} map[string]string "Invalid invoice ID"
// @Failure 401 {object} map[string]string "Unauthorized"
// @Failure 500 {object} map[string]string "Internal server error"
// @Router /invoices/{id} [delete]
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
