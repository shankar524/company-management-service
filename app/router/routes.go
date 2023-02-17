package router

import (
	"github.com/gin-gonic/gin"
	"github.com/shankar524/company-management-service/app/controllers"
)

type Router struct {
	server         *gin.Engine
	companyController controllers.CompanyControllerInterface
}

func NewRouter(server *gin.Engine, companyController controllers.CompanyControllerInterface) *Router {
	return &Router{
		server,
		companyController,
	}
}

func (r *Router) SetupRoutes() {
	basePath := r.server.Group("/api")

	basePath.GET("/health", controllers.Health)

	companyRoute := basePath.Group("/company")
	{
		companyRoute.POST("/", r.companyController.Create)
		companyRoute.GET("/:id", r.companyController.Read)
		companyRoute.PUT("/:id", r.companyController.Update)
		companyRoute.DELETE("/:id", r.companyController.Delete)
	}
}
