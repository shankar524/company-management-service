package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shankar524/company-management-service/app/controllers"
	"github.com/shankar524/company-management-service/app/middlewares"
)

type Router struct {
	server            *gin.Engine
	companyController controllers.CompanyControllerInterface
	jwtSecret         string
}

func NewRouter(server *gin.Engine, companyController controllers.CompanyControllerInterface, jwtSecret string) *Router {
	return &Router{
		server,
		companyController,
		jwtSecret,
	}
}

func (r *Router) SetupRoutes() {
	r.server.GET("/health", controllers.Health)
	basePath := r.server.Group("/api")

	companyRoute := basePath.Group("/company")
	{
		companyRoute.GET("/:id", r.companyController.Read)

		// Put mutating routes as authorized
		companyRoute.Use(middlewares.Authorize(r.jwtSecret))
		companyRoute.POST("/", r.companyController.Create)
		companyRoute.PATCH("/:id", r.companyController.Update)
		companyRoute.DELETE("/:id", r.companyController.Delete)
	}
}
