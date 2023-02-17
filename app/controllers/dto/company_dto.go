package dto

type Company struct {
	Id            string `json:"id" binding:"id"`
	Name          string `json:"name" binding:"required"`
	Description   string `json:"description" binding:"description"`
	EmployeeCount int    `json:"employeeCount" binding:"required,employee_count"`
	Registered    bool   `json:"registered" binding:"required,registered"`
	Type          string `json:"type" binding:"required,type"`
}
