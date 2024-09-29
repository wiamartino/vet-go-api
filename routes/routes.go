package routes

import (
	"go-vet/application"
	"go-vet/controllers"
	"go-vet/infrastructure/database"
	"go-vet/infrastructure/repositories"
	"go-vet/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	db, err := database.ConnectDatabase()
	if err != nil {
		panic("Failed to connect to database!")
	}

	// User
	userRepo := repositories.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{

		// Appointment
		appointmentRepo := repositories.NewAppointmentRepository(db)
		appointmentService := application.NewAppointmentService(appointmentRepo)
		appointmentController := controllers.NewAppointmentController(appointmentService)
		authorized.GET("/appointments", appointmentController.FindAppointments)
		authorized.GET("/appointments/:id", appointmentController.FindAppointment)
		authorized.POST("/appointments", appointmentController.CreateAppointment)
		authorized.PUT("/appointments/:id", appointmentController.UpdateAppointment)
		authorized.DELETE("/appointments/:id", appointmentController.DeleteAppointment)

		// Client
		clientRepo := repositories.NewClientRepository(db)
		clientService := application.NewClientService(clientRepo)
		clientController := controllers.NewClientController(clientService)
		authorized.GET("/clients", clientController.FindClients)
		authorized.GET("/clients/:id", clientController.FindClient)
		authorized.POST("/clients", clientController.CreateClient)
		authorized.PUT("/clients/:id", clientController.UpdateClient)
		authorized.DELETE("/clients/:id", clientController.DeleteClient)

		// Pet
		petRepo := repositories.NewPetRepository(db)
		petService := application.NewPetService(petRepo)
		petController := controllers.NewPetController(petService)
		authorized.GET("/pets", petController.FindPets)
		authorized.GET("/pets/:id", petController.FindPet)
		authorized.POST("/pets", petController.CreatePet)
		authorized.PUT("/pets/:id", petController.UpdatePet)
		authorized.DELETE("/pets/:id", petController.DeletePet)

		// Veterinarian
		veterinarianRepo := repositories.NewVeterinarianRepository(db)
		veterinarianService := application.NewVeterinarianService(veterinarianRepo)
		veterinarianController := controllers.NewVeterinarianController(veterinarianService)
		authorized.GET("/veterinarians", veterinarianController.FindVeterinarians)
		authorized.GET("/veterinarians/:id", veterinarianController.FindVeterinarian)
		authorized.POST("/veterinarians", veterinarianController.CreateVeterinarian)
		authorized.PUT("/veterinarians/:id", veterinarianController.UpdateVeterinarian)
		authorized.DELETE("/veterinarians/:id", veterinarianController.DeleteVeterinarian)

		// Treatment
		treatmentRepo := repositories.NewTreatmentRepository(db)
		treatmentService := application.NewTreatmentService(treatmentRepo)
		treatmentController := controllers.NewTreatmentController(treatmentService)
		authorized.GET("/treatments", treatmentController.FindTreatments)
		authorized.GET("/treatments/:id", treatmentController.FindTreatment)
		authorized.POST("/treatments", treatmentController.CreateTreatment)
		authorized.PUT("/treatments/:id", treatmentController.UpdateTreatment)
		authorized.DELETE("/treatments/:id", treatmentController.DeleteTreatment)

		// Invoice
		invoiceRepo := repositories.NewInvoiceRepository(db)
		invoiceService := application.NewInvoiceService(invoiceRepo)
		invoiceController := controllers.NewInvoiceController(invoiceService)
		authorized.GET("/invoices", invoiceController.FindInvoices)
		authorized.GET("/invoices/:id", invoiceController.FindInvoice)
		authorized.POST("/invoices", invoiceController.CreateInvoice)
		authorized.PUT("/invoices/:id", invoiceController.UpdateInvoice)
		authorized.DELETE("/invoices/:id", invoiceController.DeleteInvoice)

		// Medication
		medicationRepo := repositories.NewMedicationRepository(db)
		medicationService := application.NewMedicationService(medicationRepo)
		medicationController := controllers.NewMedicationController(medicationService)
		authorized.GET("/medications", medicationController.FindMedications)
		authorized.GET("/medications/:id", medicationController.FindMedication)
		authorized.POST("/medications", medicationController.CreateMedication)
		authorized.PUT("/medications/:id", medicationController.UpdateMedication)
		authorized.DELETE("/medications/:id", medicationController.DeleteMedication)

	}

	return r
}
