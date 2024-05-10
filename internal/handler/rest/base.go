package rest

import (
	"github.com/employee_manager/internal/service/employee"
	"github.com/gin-gonic/gin"
)

type RestHandler interface {
	ListEmployees(c *gin.Context)
	CreateEmployee(c *gin.Context)
	GetEmployeeById(c *gin.Context)
	DeleteEmployee(c *gin.Context)
	UpdateEmployee(c *gin.Context)
}

type restHandler struct {
	employeeService employee.EmployeeService
}

func NewRestHandler(employeeService employee.EmployeeService) RestHandler {
	return &restHandler{employeeService: employeeService}
}
