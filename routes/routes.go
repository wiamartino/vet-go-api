package routes

import (
	"go-vet/controllers"
	"go-vet/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{

		// Client routes
		authorized.GET("/clients", controllers.FindClients)
		authorized.POST("/clients", controllers.CreateClient)
		authorized.GET("/clients/:id", controllers.FindClient)
		authorized.PUT("/clients/:id", controllers.UpdateClient)
		authorized.DELETE("/clients/:id", controllers.DeleteClient)

		// Pet routes
		authorized.GET("/pets", controllers.FindPets)
		authorized.POST("/pets", controllers.CreatePet)
		authorized.GET("/pets/:id", controllers.FindPet)
		authorized.PUT("/pets/:id", controllers.UpdatePet)
		authorized.DELETE("/pets/:id", controllers.DeletePet)

		// Veterinarian routes
		authorized.GET("/veterinarians", controllers.FindVeterinarians)
		authorized.POST("/veterinarians", controllers.CreateVeterinarian)
		authorized.GET("/veterinarians/:id", controllers.FindVeterinarian)
		authorized.PUT("/veterinarians/:id", controllers.UpdateVeterinarian)
		authorized.DELETE("/veterinarians/:id", controllers.DeleteVeterinarian)

		// Appointment routes
		authorized.GET("/appointments", controllers.FindAppointments)
		authorized.POST("/appointments", controllers.CreateAppointment)
		authorized.GET("/appointments/:id", controllers.FindAppointment)
		authorized.PUT("/appointments/:id", controllers.UpdateAppointment)
		authorized.DELETE("/appointments/:id", controllers.DeleteAppointment)

		// Treatment routes
		authorized.GET("/treatments", controllers.FindTreatments)
		authorized.POST("/treatments", controllers.CreateTreatment)
		authorized.GET("/treatments/:id", controllers.FindTreatment)
		authorized.PUT("/treatments/:id", controllers.UpdateTreatment)
		authorized.DELETE("/treatments/:id", controllers.DeleteTreatment)

		// Invoice routes
		authorized.GET("/invoices", controllers.FindInvoices)
		authorized.POST("/invoices", controllers.CreateInvoice)
		authorized.GET("/invoices/:id", controllers.FindInvoice)
		authorized.PUT("/invoices/:id", controllers.UpdateInvoice)
		authorized.DELETE("/invoices/:id", controllers.DeleteInvoice)

		// Medication routes
		authorized.POST("/medications", controllers.CreateMedication)
		authorized.GET("/medications", controllers.FindMedications)
		authorized.GET("/medications/:id", controllers.FindMedication)
		authorized.PUT("/medications/:id", controllers.UpdateMedication)
		authorized.DELETE("/medications/:id", controllers.DeleteMedication)

	}

	return r
}
