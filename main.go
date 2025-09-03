package main

import (
	"golang-contact-management-restful-api/config"
	"golang-contact-management-restful-api/internal/database"
	"golang-contact-management-restful-api/internal/middleware"
	"golang-contact-management-restful-api/internal/server"
	userHandler "golang-contact-management-restful-api/modules/user/handler"
	userRepository "golang-contact-management-restful-api/modules/user/repository"
	userUsecase "golang-contact-management-restful-api/modules/user/usecase"

	contactHandler "golang-contact-management-restful-api/modules/contact/handler"
	contactRepository "golang-contact-management-restful-api/modules/contact/repository"
	contactUsecase "golang-contact-management-restful-api/modules/contact/usecase"

	addressHandler "golang-contact-management-restful-api/modules/address/handler"
	addressRepository "golang-contact-management-restful-api/modules/address/repository"
	addressUsecase "golang-contact-management-restful-api/modules/address/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

func main() {
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{})
	log.SetLevel(logrus.InfoLevel)

	cfg, err := config.LoadConfig()
	if err != nil {
		log.WithError(err).Fatal("Failed to load configuration file")
	}

	db := database.NewPostgresDatabase()
	err = db.Connect(cfg.Database.URL)
	if err != nil {
		log.WithError(err).Fatal("Failed to connect to database")
	}
	defer db.Close()

	srv := server.NewFiberServer(cfg)
	validate := validator.New()

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
