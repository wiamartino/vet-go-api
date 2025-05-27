package routes

import (
	"go-vet/application"
	"go-vet/controllers"
	"go-vet/infrastructure/database"
	"go-vet/infrastructure/repositories"
	"go-vet/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *database.DB) *gin.Engine {
	r := gin.Default()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

	r.Use(cors.New(config))
	r.Use(middlewares.MetricsMiddleware())

	// Public auth routes
	setupAuthRoutes(r, db)

	// Authentication & audit required for these routes
	authorized := r.Group("/api/v1")
	authorized.Use(middlewares.AuthMiddleware())
	authorized.Use(middlewares.AuditMiddleware(db))
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

func setupAuthRoutes(r *gin.Engine, db *database.DB) {
	userRepo := repositories.NewUserRepository(db)
	userService := application.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)

	auth := r.Group("/auth")
	{
		auth.POST("/register", authController.Register)
		auth.POST("/login", authController.Login)
		auth.POST("/refresh", authController.RefreshToken)
	}
}

func setupAppointmentRoutes(r *gin.RouterGroup, db *database.DB) {
	appointmentRepo := repositories.NewAppointmentRepository(db)
	appointmentService := application.NewAppointmentService(appointmentRepo)
	appointmentController := controllers.NewAppointmentController(appointmentService)

	appointments := r.Group("/appointments")
	{
		appointments.GET("", appointmentController.FindAppointments)
		appointments.GET("/:id", appointmentController.FindAppointment)
		appointments.POST("", appointmentController.CreateAppointment)
		appointments.PUT("/:id", appointmentController.UpdateAppointment)
		appointments.DELETE("/:id", appointmentController.DeleteAppointment)
	}
}

func setupClientRoutes(r *gin.RouterGroup, db *database.DB) {
	clientRepo := repositories.NewClientRepository(db)
	clientService := application.NewClientService(clientRepo)
	clientController := controllers.NewClientController(clientService)

	clients := r.Group("/clients")
	{
		clients.GET("", clientController.FindClients)
		clients.GET("/:id", clientController.FindClient)
		clients.POST("", clientController.CreateClient)
		clients.PUT("/:id", clientController.UpdateClient)
		clients.DELETE("/:id", clientController.DeleteClient)
	}
}

func setupPetRoutes(r *gin.RouterGroup, db *database.DB) {
	petRepo := repositories.NewPetRepository(db)
	petService := application.NewPetService(petRepo)
	petController := controllers.NewPetController(petService)

	pets := r.Group("/pets")
	{
		pets.GET("", petController.FindPets)
		pets.GET("/:id", petController.FindPet)
		pets.POST("", petController.CreatePet)
		pets.PUT("/:id", petController.UpdatePet)
		pets.DELETE("/:id", petController.DeletePet)
	}
}

func setupVeterinarianRoutes(r *gin.RouterGroup, db *database.DB) {
	veterinarianRepo := repositories.NewVeterinarianRepository(db)
	veterinarianService := application.NewVeterinarianService(veterinarianRepo)
	veterinarianController := controllers.NewVeterinarianController(veterinarianService)

	veterinarians := r.Group("/veterinarians")
	{
		veterinarians.GET("", veterinarianController.FindVeterinarians)
		veterinarians.GET("/:id", veterinarianController.FindVeterinarian)
		veterinarians.POST("", veterinarianController.CreateVeterinarian)
		veterinarians.PUT("/:id", veterinarianController.UpdateVeterinarian)
		veterinarians.DELETE("/:id", veterinarianController.DeleteVeterinarian)
	}
}

func setupTreatmentRoutes(r *gin.RouterGroup, db *database.DB) {
	treatmentRepo := repositories.NewTreatmentRepository(db)
	treatmentService := application.NewTreatmentService(treatmentRepo)
	treatmentController := controllers.NewTreatmentController(treatmentService)

	treatments := r.Group("/treatments")
	{
		treatments.GET("", treatmentController.FindTreatments)
		treatments.GET("/:id", treatmentController.FindTreatment)
		treatments.POST("", treatmentController.CreateTreatment)
		treatments.PUT("/:id", treatmentController.UpdateTreatment)
		treatments.DELETE("/:id", treatmentController.DeleteTreatment)
	}
}

func setupInvoiceRoutes(r *gin.RouterGroup, db *database.DB) {
	invoiceRepo := repositories.NewInvoiceRepository(db)
	invoiceService := application.NewInvoiceService(invoiceRepo)
	invoiceController := controllers.NewInvoiceController(invoiceService)

	invoices := r.Group("/invoices")
	{
		invoices.GET("", invoiceController.FindInvoices)
		invoices.GET("/:id", invoiceController.FindInvoice)
		invoices.POST("", invoiceController.CreateInvoice)
		invoices.PUT("/:id", invoiceController.UpdateInvoice)
		invoices.DELETE("/:id", invoiceController.DeleteInvoice)
	}
}

func setupMedicationRoutes(r *gin.RouterGroup, db *database.DB) {
	medicationRepo := repositories.NewMedicationRepository(db)
	medicationService := application.NewMedicationService(medicationRepo)
	medicationController := controllers.NewMedicationController(medicationService)

	medications := r.Group("/medications")
	{
		medications.GET("", medicationController.FindMedications)
		medications.GET("/:id", medicationController.FindMedication)
		medications.POST("", medicationController.CreateMedication)
		medications.PUT("/:id", medicationController.UpdateMedication)
		medications.DELETE("/:id", medicationController.DeleteMedication)
	}
}
