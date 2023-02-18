package dto

type UpdateCompanyRequest struct {
	Name          string `json:"name" binding:"min=3,max=15"`
	Description   string `json:"description" binding:"max=3000"`
	EmployeeCount int    `json:"employeeCount" binding:"gte=1"`
	Registered    bool   `json:"registered"`
	Type          string `json:"type" binding:"oneof=Corporations NonProfit Cooperative Sole\\ Proprietorship"`
}
