package employee

import (
	"github.com/employee_manager/internal/dao"
	"github.com/employee_manager/internal/repo/employee"
)

type EmployeeService interface {
	ListEmployees(page int, pageSize int) ([]*dao.Employee, error)
	CreateEmployee(employee *dao.Employee) (*dao.Employee, error)
	GetEmployeeById(id int) (*dao.Employee, error)
	UpdateEmployee(employee *dao.Employee) error
	DeleteEmployee(id int) error
}

type employeeService struct {
	employeeRepo employee.EmployeeRepo
}

func NewEmployeeService(employeeRepo employee.EmployeeRepo) EmployeeService {
	return &employeeService{employeeRepo: employeeRepo}
}
