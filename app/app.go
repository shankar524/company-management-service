package app

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	pg "github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/shankar524/company-management-service/app/configs"
	"github.com/shankar524/company-management-service/app/controllers"
	"github.com/shankar524/company-management-service/app/db"
	"github.com/shankar524/company-management-service/app/messageBroker"
	"github.com/shankar524/company-management-service/app/migrations"
	"github.com/shankar524/company-management-service/app/repositories"
	"github.com/shankar524/company-management-service/app/router"
	"github.com/shankar524/company-management-service/app/services"
)

func SetupEngine(config configs.Config) *gin.Engine {
	if config.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}

	engine := gin.Default()

	/*
		====== Setup Postgres ============
	*/
	db := db.NewPgDB(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", config.Database.Host, config.Database.Port),
		User:     config.Database.User,
		Password: config.Database.Password,
		Database: config.Database.Name,
	})

	migrations.Init(db)

	/*
		====== Setup Kafka ============
	*/
	kafkaHost := fmt.Sprintf("%s:%s", config.Kafka.Host, config.Kafka.Port)
	kafkaClient := messageBroker.NewKafkaClient(kafkaHost)

	messageService := services.NewMessageService(kafkaClient)
	companyRepo := repositories.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo, messageService)
	companyController := controllers.NewCompanyController(companyService)

	routes := router.NewRouter(engine, companyController, config.JWTSecret)

	/*
		====== Setup Controllers and Routers ============
	*/
	routes.SetupRoutes()

	return engine
}

func Run() {
	/*
		====== Setup configs ============
	*/
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	config := configs.GetConfig()
	engine := SetupEngine(config)
	engine.Run(fmt.Sprintf(":%s", config.Port))
}
