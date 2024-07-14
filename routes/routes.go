package routes

import (
	"go-vet/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Client routes
	r.GET("/clients", controllers.FindClients)
	r.POST("/clients", controllers.CreateClient)
	r.GET("/clients/:id", controllers.FindClient)
	r.PUT("/clients/:id", controllers.UpdateClient)
	r.DELETE("/clients/:id", controllers.DeleteClient)

	// Pet routes
	r.GET("/pets", controllers.FindPets)
	r.POST("/pets", controllers.CreatePet)
	r.GET("/pets/:id", controllers.FindPet)
	r.PUT("/pets/:id", controllers.UpdatePet)
	r.DELETE("/pets/:id", controllers.DeletePet)

	// Veterinarian routes
	r.GET("/veterinarians", controllers.FindVeterinarians)
	r.POST("/veterinarians", controllers.CreateVeterinarian)
	r.GET("/veterinarians/:id", controllers.FindVeterinarian)
	r.PUT("/veterinarians/:id", controllers.UpdateVeterinarian)
	r.DELETE("/veterinarians/:id", controllers.DeleteVeterinarian)

	// Appointment routes
	r.GET("/appointments", controllers.FindAppointments)
	r.POST("/appointments", controllers.CreateAppointment)
	r.GET("/appointments/:id", controllers.FindAppointment)
	r.PUT("/appointments/:id", controllers.UpdateAppointment)
	r.DELETE("/appointments/:id", controllers.DeleteAppointment)

	// Treatment routes
	r.GET("/treatments", controllers.FindTreatments)
	r.POST("/treatments", controllers.CreateTreatment)
	r.GET("/treatments/:id", controllers.FindTreatment)
	r.PUT("/treatments/:id", controllers.UpdateTreatment)
	r.DELETE("/treatments/:id", controllers.DeleteTreatment)

	// Invoice routes
	r.GET("/invoices", controllers.FindInvoices)
	r.POST("/invoices", controllers.CreateInvoice)
	r.GET("/invoices/:id", controllers.FindInvoice)
	r.PUT("/invoices/:id", controllers.UpdateInvoice)
	r.DELETE("/invoices/:id", controllers.DeleteInvoice)

	return r
}
