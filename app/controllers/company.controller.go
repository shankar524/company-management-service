package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shankar524/company-management-service/app/controllers/dto"
	"github.com/shankar524/company-management-service/app/services"
)

type CompanyControllerInterface interface {
	Create(c *gin.Context)
	Read(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type CompanyController struct {
	companyService services.CompanyServiceInterface
}

func NewCompanyController(service services.CompanyServiceInterface) CompanyControllerInterface {
	return &CompanyController{service}
}

func (ctr *CompanyController) Create(c *gin.Context) {
	u := dto.CreateCompanyRequest{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	id, err := ctr.companyService.CreateCompany(context.Background(), u)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Printf("Request received: %+v \n", u)
	c.JSON(http.StatusOK, gin.H{"message": "created", "id": id})
}

func (ctr *CompanyController) Read(c *gin.Context) {
	e := c.Param("id")
	fmt.Printf("Path param received: %+v \n", e)

	u, err := ctr.companyService.GetCompany(context.Background(), e)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, u)
}

func (ctr *CompanyController) Update(c *gin.Context) {

	u := dto.UpdateCompanyRequest{}
	if err := c.ShouldBindJSON(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	e := c.Param("id")

	fmt.Printf("Request received: %+v \n", u)
	fmt.Printf("Path param received: %+v \n", e)

	if err := ctr.companyService.UpdateCompany(context.Background(), u, e); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}

func (ctr *CompanyController) Delete(c *gin.Context) {
	id := c.Param("id")

	fmt.Printf("Path param received: %+v \n", id)

	if err := ctr.companyService.DeleteCompany(context.Background(), id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
