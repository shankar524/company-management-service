package app

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	// "github.com/go-pg/pg"
	pg "github.com/go-pg/pg/v10"
	"github.com/joho/godotenv"
	"github.com/shankar524/company-management-service/app/configs"
	"github.com/shankar524/company-management-service/app/controllers"
	"github.com/shankar524/company-management-service/app/db"
	"github.com/shankar524/company-management-service/app/migrations"
	"github.com/shankar524/company-management-service/app/repositories"
	"github.com/shankar524/company-management-service/app/router"
	"github.com/shankar524/company-management-service/app/services"
)

var (
	engine = gin.Default()
	config configs.Config
)

func Run() {
	/*
		====== Setup configs ============
	*/
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	} else {
		log.Println("ENV file loaded successfully")
		log.Print(os.Getenv("DB_HOST"))
	}
	config := configs.GetConfig()

	/*
		====== Setup DB Connections and migrations ============
	*/
	db := db.NewPgDB(&pg.Options{
		Addr:     fmt.Sprintf("%s:%s", "host.docker.internal", config.Database.Port),
		User:     config.Database.User,
		Password: config.Database.Password,
		Database: config.Database.Name,
	})

	migrations.Init(db)

	/*
		====== Setup Controllers and Routers ============
	*/

	companyRepo := repositories.NewCompanyRepository(db)
	companyService := services.NewCompanyService(companyRepo)
	companyController := controllers.NewCompanyController(companyService)

	routes := router.NewRouter(engine, companyController)
	routes.SetupRoutes()

	engine.Run(fmt.Sprintf(":%s", config.Port))
}
