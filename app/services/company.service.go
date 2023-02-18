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

const (
	Topic = "Company"
)

type CompanyService struct {
	cRepo            repo.CompanyRepositoryInterface
	messagePublisher Broadcaster
}

func NewCompanyService(cRepo repo.CompanyRepositoryInterface, messagePublisher Broadcaster) CompanyServiceInterface {
	return &CompanyService{cRepo, messagePublisher}
}

func (c *CompanyService) CreateCompany(ctx context.Context, request dto.CreateCompanyRequest) (id string, err error) {
	id, err = c.cRepo.Create(ctx, request)
	if err != nil {
		return
	}

	event := JSON{
		"id":    id,
		"event": "CREATE",
	}
	err = c.messagePublisher.Broadcast(Topic, event)
	return
}

func (c *CompanyService) GetCompany(ctx context.Context, id string) (*dto.Company, error) {
	return c.cRepo.GetById(ctx, id)
}

func (c *CompanyService) UpdateCompany(ctx context.Context, request dto.UpdateCompanyRequest, id string) (err error) {

	err = c.cRepo.UpdateById(ctx, request, id)
	if err != nil {
		return
	}

	event := JSON{
		"id":     id,
		"action": "UPDATE",
	}
	err = c.messagePublisher.Broadcast(Topic, event)
	return
}

func (c *CompanyService) DeleteCompany(ctx context.Context, id string) (err error) {

	err = c.cRepo.DeleteById(ctx, id)
	if err != nil {
		return
	}

	event := JSON{
		"id":    id,
		"event": "DELETE",
	}
	err = c.messagePublisher.Broadcast(Topic, event)
	return
}
