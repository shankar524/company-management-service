package services

import (
	"context"

	"github.com/shankar524/company-management-service/app/controllers/dto"

	repo "github.com/shankar524/company-management-service/app/repositories"
)

type CompanyServiceInterface interface {
	CreateCompany(ctx context.Context, request dto.CreateCompanyRequest) (string, error)
	GetCompany(ctx context.Context, id string) (*dto.Company, error)
	UpdateCompany(ctx context.Context, request dto.UpdateCompanyRequest, id string) error
	DeleteCompany(ctx context.Context, id string) error
}

type CompanyService struct {
	cRepo repo.CompanyRepositoryInterface
}

func NewCompanyService(cRepo repo.CompanyRepositoryInterface) CompanyServiceInterface {
	return &CompanyService{cRepo}
}

func (s *CompanyService) CreateCompany(ctx context.Context, request dto.CreateCompanyRequest) (string, error) {
	return s.cRepo.Create(ctx, request)
}

func (c *CompanyService) GetCompany(ctx context.Context, id string) (*dto.Company, error) {
	return c.cRepo.GetById(ctx, id)
}

func (c *CompanyService) UpdateCompany(ctx context.Context, request dto.UpdateCompanyRequest, id string) error {

	return c.cRepo.UpdateById(ctx, request, id)
}

func (c *CompanyService) DeleteCompany(ctx context.Context, id string) error {

	return c.cRepo.DeleteById(ctx, id)
}
