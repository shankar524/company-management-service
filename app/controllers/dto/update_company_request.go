package dto

type UpdateCompanyRequest struct {
	Name          string `json:"name"`
	Description   string `json:"description"`
	EmployeeCount int    `json:"employeeCount"`
	Registered    bool   `json:"registered"`
	Type          string `json:"type"`
}
