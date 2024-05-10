package employee

import (
	"github.com/employee_manager/internal/dao"
	"sync"
)

type EmployeeRepo interface {
	CreateEmployee(data *dao.Employee) (*dao.Employee, error)
	GetEmployeeById(id int) (*dao.Employee, error)
	UpdateEmployee(data *dao.Employee) error
	DeleteEmployee(data *dao.Employee) error
	ListEmployees(page int, pageSize int) ([]*dao.Employee, error)
}

type employeeRepo struct {
	inMemoryStore map[int]*dao.Employee
	mu            sync.RWMutex
	idCounter     int
}

func NewEmployeeRepo() EmployeeRepo {
	return &employeeRepo{
		inMemoryStore: make(map[int]*dao.Employee),
		mu:            sync.RWMutex{},
		idCounter:     0,
	}
}
