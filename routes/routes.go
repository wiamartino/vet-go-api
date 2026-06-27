package routes

import (
	"net/http"
	"os"

	"go-vet/application"
	"go-vet/controllers"
	"go-vet/infrastructure/database"
	"go-vet/infrastructure/repositories"
	"go-vet/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRouter(db *database.DB) *gin.Engine {
	r := gin.Default()

	// Use the provided database connection instead of creating a new one
	if db == nil {
		panic("Database connection is required!")
	}

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowCredentials = true
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}

	r.Use(cors.New(config))
	r.Use(middlewares.MetricsMiddleware())

	// Swagger documentation route
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Health check endpoint
	r.GET("/healthz", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "database unavailable"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "OK"})
	})

	// Public auth routes
	setupAuthRoutes(r, db)

	// Log authentication status
	if os.Getenv("DISABLE_AUTH") == "true" {
		logrus.Warn("⚠️⚠️⚠️  AUTHENTICATION DISABLED - USE ONLY IN DEVELOPMENT  ⚠️⚠️⚠️")
	} else {
		logrus.Info("✓ Authentication enabled")
	}

	// Authentication & audit required for these routes
	authorized := r.Group("/api/v1")
	authorized.Use(middlewares.AuthMiddleware())
	authorized.Use(middlewares.AuditMiddleware(db))
	{
		// Metrics endpoint
		authorized.GET("/metrics", middlewares.RoleAuthMiddleware("admin"), func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"status": "success",
				"data": gin.H{
					"total_requests": middlewares.GetTotalRequests(),
					"total_errors":   middlewares.GetTotalErrors(),
				},
			})
		})

		setupAppointmentRoutes(authorized, db)
		setupClientRoutes(authorized, db)
		setupPetRoutes(authorized, db)
		setupVeterinarianRoutes(authorized, db)
		setupTreatmentRoutes(authorized, db)
		setupInvoiceRoutes(authorized, db)
		setupMedicationRoutes(authorized, db)
		setupMedicalRecordRoutes(authorized, db)
		setupVaccinationRoutes(authorized, db)
		setupSurgeryRoutes(authorized, db)
		setupAllergyRoutes(authorized, db)
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
		appointments.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), appointmentController.FindAppointments)
		appointments.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), appointmentController.FindAppointment)
		appointments.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), appointmentController.CreateAppointment)
		appointments.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), appointmentController.UpdateAppointment)
		appointments.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), appointmentController.DeleteAppointment)
	}
}

func setupClientRoutes(r *gin.RouterGroup, db *database.DB) {
	clientRepo := repositories.NewClientRepository(db)
	clientService := application.NewClientService(clientRepo)
	clientController := controllers.NewClientController(clientService)

	clients := r.Group("/clients")
	{
		clients.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), clientController.FindClients)
		clients.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), clientController.FindClient)
		clients.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), clientController.CreateClient)
		clients.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), clientController.UpdateClient)
		clients.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), clientController.DeleteClient)
	}
}

func setupPetRoutes(r *gin.RouterGroup, db *database.DB) {
	petRepo := repositories.NewPetRepository(db)
	petService := application.NewPetService(petRepo)
	petController := controllers.NewPetController(petService)

	// Medical records for pets
	medicalRecordRepo := repositories.NewMedicalRecordRepository(db)
	medicalRecordService := application.NewMedicalRecordService(medicalRecordRepo)
	medicalRecordController := controllers.NewMedicalRecordController(medicalRecordService)

	// Vaccinations for pets
	vaccinationRepo := repositories.NewVaccinationRepository(db)
	vaccinationService := application.NewVaccinationService(vaccinationRepo)
	vaccinationController := controllers.NewVaccinationController(vaccinationService)

	// Surgeries for pets
	surgeryRepo := repositories.NewSurgeryRepository(db)
	surgeryService := application.NewSurgeryService(surgeryRepo)
	surgeryController := controllers.NewSurgeryController(surgeryService)

	// Allergies for pets
	allergyRepo := repositories.NewAllergyRepository(db)
	allergyService := application.NewAllergyService(allergyRepo)
	allergyController := controllers.NewAllergyController(allergyService)

	pets := r.Group("/pets")
	{
		pets.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), petController.FindPets)
		pets.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), petController.FindPet)
		pets.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), petController.CreatePet)
		pets.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), petController.UpdatePet)
		pets.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), petController.DeletePet)

		// Nested resources - using :id for consistency
		pets.GET("/:id/medical-records", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicalRecordController.GetMedicalRecordsByPet)
		pets.GET("/:id/vaccinations", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), vaccinationController.GetVaccinationsByPet)
		pets.GET("/:id/surgeries", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.GetSurgeriesByPet)
		pets.GET("/:id/allergies", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), allergyController.GetAllergiesByPet)
		pets.GET("/:id/allergies/active", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), allergyController.GetActiveAllergiesByPet)
	}
}

func setupVeterinarianRoutes(r *gin.RouterGroup, db *database.DB) {
	veterinarianRepo := repositories.NewVeterinarianRepository(db)
	veterinarianService := application.NewVeterinarianService(veterinarianRepo)
	veterinarianController := controllers.NewVeterinarianController(veterinarianService)

	veterinarians := r.Group("/veterinarians")
	{
		veterinarians.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), veterinarianController.FindVeterinarians)
		veterinarians.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), veterinarianController.FindVeterinarian)
		veterinarians.POST("", middlewares.RoleAuthMiddleware("admin"), veterinarianController.CreateVeterinarian)
		veterinarians.PUT("/:id", middlewares.RoleAuthMiddleware("admin"), veterinarianController.UpdateVeterinarian)
		veterinarians.DELETE("/:id", middlewares.RoleAuthMiddleware("admin"), veterinarianController.DeleteVeterinarian)
	}
}

func setupTreatmentRoutes(r *gin.RouterGroup, db *database.DB) {
	treatmentRepo := repositories.NewTreatmentRepository(db)
	treatmentService := application.NewTreatmentService(treatmentRepo)
	treatmentController := controllers.NewTreatmentController(treatmentService)

	treatments := r.Group("/treatments")
	{
		treatments.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), treatmentController.FindTreatments)
		treatments.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), treatmentController.FindTreatment)
		treatments.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), treatmentController.CreateTreatment)
		treatments.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), treatmentController.UpdateTreatment)
		treatments.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), treatmentController.DeleteTreatment)
	}
}

func setupInvoiceRoutes(r *gin.RouterGroup, db *database.DB) {
	invoiceRepo := repositories.NewInvoiceRepository(db)
	invoiceService := application.NewInvoiceService(invoiceRepo)
	invoiceController := controllers.NewInvoiceController(invoiceService)

	invoices := r.Group("/invoices")
	{
		invoices.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), invoiceController.FindInvoices)
		invoices.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), invoiceController.FindInvoice)
		invoices.POST("", middlewares.RoleAuthMiddleware("admin"), invoiceController.CreateInvoice)
		invoices.PUT("/:id", middlewares.RoleAuthMiddleware("admin"), invoiceController.UpdateInvoice)
		invoices.DELETE("/:id", middlewares.RoleAuthMiddleware("admin"), invoiceController.DeleteInvoice)
	}
}

func setupMedicationRoutes(r *gin.RouterGroup, db *database.DB) {
	medicationRepo := repositories.NewMedicationRepository(db)
	medicationService := application.NewMedicationService(medicationRepo)
	medicationController := controllers.NewMedicationController(medicationService)

	medications := r.Group("/medications")
	{
		medications.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), medicationController.FindMedications)
		medications.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), medicationController.FindMedication)
		medications.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicationController.CreateMedication)
		medications.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicationController.UpdateMedication)
		medications.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicationController.DeleteMedication)
	}
}

func setupMedicalRecordRoutes(r *gin.RouterGroup, db *database.DB) {
	medicalRecordRepo := repositories.NewMedicalRecordRepository(db)
	medicalRecordService := application.NewMedicalRecordService(medicalRecordRepo)
	medicalRecordController := controllers.NewMedicalRecordController(medicalRecordService)

	medicalRecords := r.Group("/medical-records")
	{
		medicalRecords.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicalRecordController.GetMedicalRecords)
		medicalRecords.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicalRecordController.GetMedicalRecord)
		medicalRecords.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicalRecordController.CreateMedicalRecord)
		medicalRecords.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicalRecordController.UpdateMedicalRecord)
		medicalRecords.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), medicalRecordController.DeleteMedicalRecord)
	}
}

func setupVaccinationRoutes(r *gin.RouterGroup, db *database.DB) {
	vaccinationRepo := repositories.NewVaccinationRepository(db)
	vaccinationService := application.NewVaccinationService(vaccinationRepo)
	vaccinationController := controllers.NewVaccinationController(vaccinationService)

	vaccinations := r.Group("/vaccinations")
	{
		vaccinations.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.GetVaccinations)
		vaccinations.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.GetVaccination)
		vaccinations.GET("/due", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.GetDueVaccinations)
		vaccinations.GET("/overdue", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.GetOverdueVaccinations)
		vaccinations.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.CreateVaccination)
		vaccinations.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.UpdateVaccination)
		vaccinations.PATCH("/:id/complete", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.CompleteVaccination)
		vaccinations.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), vaccinationController.DeleteVaccination)
	}
}

func setupSurgeryRoutes(r *gin.RouterGroup, db *database.DB) {
	surgeryRepo := repositories.NewSurgeryRepository(db)
	surgeryService := application.NewSurgeryService(surgeryRepo)
	surgeryController := controllers.NewSurgeryController(surgeryService)

	surgeries := r.Group("/surgeries")
	{
		surgeries.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.GetSurgeries)
		surgeries.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.GetSurgery)
		surgeries.GET("/status", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.GetSurgeriesByStatus)
		surgeries.GET("/scheduled", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.GetScheduledSurgeries)
		surgeries.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.CreateSurgery)
		surgeries.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.UpdateSurgery)
		surgeries.PATCH("/:id/start", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.StartSurgery)
		surgeries.PATCH("/:id/complete", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.CompleteSurgery)
		surgeries.PATCH("/:id/cancel", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.CancelSurgery)
		surgeries.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), surgeryController.DeleteSurgery)
	}
}

func setupAllergyRoutes(r *gin.RouterGroup, db *database.DB) {
	allergyRepo := repositories.NewAllergyRepository(db)
	allergyService := application.NewAllergyService(allergyRepo)
	allergyController := controllers.NewAllergyController(allergyService)

	allergies := r.Group("/allergies")
	{
		allergies.GET("", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), allergyController.GetAllergies)
		allergies.GET("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), allergyController.GetAllergy)
		allergies.GET("/severity", middlewares.RoleAuthMiddleware("admin", "veterinarian", "user"), allergyController.GetAllergiesBySeverity)
		allergies.POST("", middlewares.RoleAuthMiddleware("admin", "veterinarian"), allergyController.CreateAllergy)
		allergies.PUT("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), allergyController.UpdateAllergy)
		allergies.PATCH("/:id/deactivate", middlewares.RoleAuthMiddleware("admin", "veterinarian"), allergyController.DeactivateAllergy)
		allergies.PATCH("/:id/reactivate", middlewares.RoleAuthMiddleware("admin", "veterinarian"), allergyController.ReactivateAllergy)
		allergies.DELETE("/:id", middlewares.RoleAuthMiddleware("admin", "veterinarian"), allergyController.DeleteAllergy)
	}
}
