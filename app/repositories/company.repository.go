package repositories

import (
	"context"

	"github.com/go-pg/pg/v10"
	"github.com/shankar524/company-management-service/app/controllers/dto"
)

type CompanyRepositoryInterface interface {
	Create(ctx context.Context, request dto.CreateCompanyRequest) (string, error)
	GetById(ctx context.Context, id string) (*dto.Company, error)
	UpdateById(ctx context.Context, request dto.UpdateCompanyRequest, id string) error
	DeleteById(ctx context.Context, id string) error
}

type CompanyRepository struct {
	db *pg.DB
}

func NewCompanyRepository(db *pg.DB) CompanyRepositoryInterface {
	return &CompanyRepository{db}
}

func (r *CompanyRepository) Create(ctx context.Context, request dto.CreateCompanyRequest) (id string, err error) {
	company := dto.Company{
		Name:          request.Name,
		Description:   request.Description,
		EmployeeCount: request.EmployeeCount,
		Registered:    request.Registered,
		Type:          request.Type,
	}
	if _, err = r.db.Model(&company).Context(ctx).Insert(); err == nil {
		id = company.Id
	}

	return id, err
}

func (r *CompanyRepository) GetById(ctx context.Context, id string) (*dto.Company, error) {
	company := dto.Company{}
	err := r.db.Model(&company).Context(ctx).Where("id = ?", id).Select()

	return &company, err
}

func (r *CompanyRepository) UpdateById(ctx context.Context, request dto.UpdateCompanyRequest, id string) error {
	company := dto.Company{}
	_, err := r.db.Model(&company).Context(ctx).Set("name = ?, description = ?, employee_count = ?, registered = ?, type = ?", request.Name, request.Description, request.EmployeeCount, request.Registered, request.Type).Where("id = ?", id).Update()

	return err
}

func (r *CompanyRepository) DeleteById(ctx context.Context, id string) error {
	company := dto.Company{}
	_, err := r.db.Model(&company).Context(ctx).Where("id = ?", id).Delete()

	return err
}
