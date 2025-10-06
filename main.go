package main

import (
	"strings"

	"golang-contact-management-restful-api/config"
	"golang-contact-management-restful-api/internal/database"
	"golang-contact-management-restful-api/internal/middleware"
	"golang-contact-management-restful-api/internal/server"
	addressHandler "golang-contact-management-restful-api/modules/address/handler"
	addressRepository "golang-contact-management-restful-api/modules/address/repository"
	addressUsecase "golang-contact-management-restful-api/modules/address/usecase"
	contactHandler "golang-contact-management-restful-api/modules/contact/handler"
	contactRepository "golang-contact-management-restful-api/modules/contact/repository"
	contactUsecase "golang-contact-management-restful-api/modules/contact/usecase"
	userHandler "golang-contact-management-restful-api/modules/user/handler"
	userRepository "golang-contact-management-restful-api/modules/user/repository"
	userUsecase "golang-contact-management-restful-api/modules/user/usecase"

	addressEntity "golang-contact-management-restful-api/modules/address/entities"
	contactEntity "golang-contact-management-restful-api/modules/contact/entities"
	userEntity "golang-contact-management-restful-api/modules/user/entities"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to load configuration")
	}

	db := database.NewPostgresDatabase()
	if err := db.Connect(cfg.Database.URL); err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}
	defer db.Close()

	log.Info("Running AutoMigrate...")
	if err := db.Gorm.AutoMigrate(
		&userEntity.User{},
		&contactEntity.Contact{},
		&addressEntity.Address{},
	); err != nil {
		log.WithError(err).Fatal("Failed to run auto migration")
	}
	log.Info("AutoMigrate complete (safe to run multiple times)")

	if err != nil {
		log.WithError(err).Fatal("Failed to run auto migration")
	}

	log.Info("Auto migration completed successfully")

	srv := server.NewFiberServer(cfg)
	validate := validator.New()

	var allowOrigins string
	if cfg.AppEnv == "production" {
		allowOrigins = cfg.Frontend.Prod
	} else {
		var origins []string
		if cfg.Frontend.Dev != "" {
			origins = append(origins, cfg.Frontend.Dev)
		}
		if cfg.Frontend.Dev2 != "" {
			origins = append(origins, cfg.Frontend.Dev2)
		}
		allowOrigins = strings.Join(origins, ", ")
	}

	srv.GetEngine().Use(cors.New(cors.Config{
		AllowOrigins:     allowOrigins,
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH",
		AllowCredentials: true,
	}))

	uRepo := userRepository.NewUserRepository(db.Gorm)
	uUC := userUsecase.NewUserUsecase(uRepo, validate)
	uH := userHandler.NewUserHttpHandler(srv.GetEngine(), uUC)

	cRepo := contactRepository.NewContactRepository(db.Gorm)
	cUC := contactUsecase.NewContactUsecase(cRepo, validate)
	cH := contactHandler.NewContactHttpHandler(srv.GetEngine(), cUC)

	hRepo := addressRepository.NewAddressRepository(db.Gorm)
	hUC := addressUsecase.NewAddressUsecase(hRepo, validate)
	hH := addressHandler.NewAddressHttpHandler(srv.GetEngine(), hUC)

	auth := middleware.RequireAuth(uRepo)

	server.RegisterUserRoutes(srv.GetEngine(), uH, auth)
	server.RegisterContactRoutes(srv.GetEngine(), cH, auth)
	server.RegisterAddressRoutes(srv.GetEngine(), hH, auth)

	log.WithField("port", cfg.Server.Port).Info("Server is running")
	if err := srv.Start(); err != nil {
		log.WithError(err).Fatal("Failed to start server")
	}
}
