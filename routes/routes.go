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
	setupUserRoutes(r, db)

	r.Use(middlewares.MetricsMiddleware())
	r.Use(middlewares.AuditMiddleware(db))

	authorized := r.Group("/")
	authorized.Use(middlewares.AuthMiddleware())
	{
		setupAppointmentRoutes(authorized, db)
		setupClientRoutes(authorized, db)
		setupPetRoutes(authorized, db)
		setupVeterinarianRoutes(authorized, db)
		setupTreatmentRoutes(authorized, db)
		setupInvoiceRoutes(authorized, db)
		setupMedicationRoutes(authorized, db)
	}

	return r
}

func setupUserRoutes(r *gin.Engine, db *database.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)

	r.POST("/register", authController.Register)
	r.POST("/login", authController.Login)
}

func setupAppointmentRoutes(r *gin.RouterGroup, db *database.DB) {
	appointmentRepo := repositories.NewAppointmentRepository(db)
	appointmentService := application.NewAppointmentService(appointmentRepo)
	appointmentController := controllers.NewAppointmentController(appointmentService)

	r.GET("/appointments", appointmentController.FindAppointments)
	r.GET("/appointments/:id", appointmentController.FindAppointment)
	r.POST("/appointments", appointmentController.CreateAppointment)
	r.PUT("/appointments/:id", appointmentController.UpdateAppointment)
	r.DELETE("/appointments/:id", appointmentController.DeleteAppointment)
}

func setupClientRoutes(r *gin.RouterGroup, db *database.DB) {
	clientRepo := repositories.NewClientRepository(db)
	clientService := application.NewClientService(clientRepo)
	clientController := controllers.NewClientController(clientService)

	r.GET("/clients", clientController.FindClients)
	r.GET("/clients/:id", clientController.FindClient)
	r.POST("/clients", clientController.CreateClient)
	r.PUT("/clients/:id", clientController.UpdateClient)
	r.DELETE("/clients/:id", clientController.DeleteClient)
}

func setupPetRoutes(r *gin.RouterGroup, db *database.DB) {
	petRepo := repositories.NewPetRepository(db)
	petService := application.NewPetService(petRepo)
	petController := controllers.NewPetController(petService)

	r.GET("/pets", petController.FindPets)
	r.GET("/pets/:id", petController.FindPet)
	r.POST("/pets", petController.CreatePet)
	r.PUT("/pets/:id", petController.UpdatePet)
	r.DELETE("/pets/:id", petController.DeletePet)
}

func setupVeterinarianRoutes(r *gin.RouterGroup, db *database.DB) {
	veterinarianRepo := repositories.NewVeterinarianRepository(db)
	veterinarianService := application.NewVeterinarianService(veterinarianRepo)
	veterinarianController := controllers.NewVeterinarianController(veterinarianService)

	r.GET("/veterinarians", veterinarianController.FindVeterinarians)
	r.GET("/veterinarians/:id", veterinarianController.FindVeterinarian)
	r.POST("/veterinarians", veterinarianController.CreateVeterinarian)
	r.PUT("/veterinarians/:id", veterinarianController.UpdateVeterinarian)
	r.DELETE("/veterinarians/:id", veterinarianController.DeleteVeterinarian)
}

func setupTreatmentRoutes(r *gin.RouterGroup, db *database.DB) {
	treatmentRepo := repositories.NewTreatmentRepository(db)
	treatmentService := application.NewTreatmentService(treatmentRepo)
	treatmentController := controllers.NewTreatmentController(treatmentService)

	r.GET("/treatments", treatmentController.FindTreatments)
	r.GET("/treatments/:id", treatmentController.FindTreatment)
	r.POST("/treatments", treatmentController.CreateTreatment)
	r.PUT("/treatments/:id", treatmentController.UpdateTreatment)
	r.DELETE("/treatments/:id", treatmentController.DeleteTreatment)
}

func setupInvoiceRoutes(r *gin.RouterGroup, db *database.DB) {
	invoiceRepo := repositories.NewInvoiceRepository(db)
	invoiceService := application.NewInvoiceService(invoiceRepo)
	invoiceController := controllers.NewInvoiceController(invoiceService)

	r.GET("/invoices", invoiceController.FindInvoices)
	r.GET("/invoices/:id", invoiceController.FindInvoice)
	r.POST("/invoices", invoiceController.CreateInvoice)
	r.PUT("/invoices/:id", invoiceController.UpdateInvoice)
	r.DELETE("/invoices/:id", invoiceController.DeleteInvoice)
}

func setupMedicationRoutes(r *gin.RouterGroup, db *database.DB) {
	medicationRepo := repositories.NewMedicationRepository(db)
	medicationService := application.NewMedicationService(medicationRepo)
	medicationController := controllers.NewMedicationController(medicationService)

	r.GET("/medications", medicationController.FindMedications)
	r.GET("/medications/:id", medicationController.FindMedication)
	r.POST("/medications", medicationController.CreateMedication)
	r.PUT("/medications/:id", medicationController.UpdateMedication)
	r.DELETE("/medications/:id", medicationController.DeleteMedication)
}
