package dto

type CreateCompanyRequest struct {
	Name          string `json:"name" binding:"required,min=3,max=15"`
	Description   string `json:"description" binding:"max=3000"`
	EmployeeCount int    `json:"employeeCount" binding:"required,gte=1"`
	Registered    bool   `json:"registered" binding:"required"`
	Type          string `json:"type" binding:"required,oneof=Corporations NonProfit Cooperative Sole\\ Proprietorship"`
}
