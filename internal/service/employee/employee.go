package employee

import (
	"fmt"
	"github.com/employee_manager/internal/common/constants"
	"github.com/employee_manager/internal/dao"
)

func (e *employeeService) ListEmployees(page int, pageSize int) ([]*dao.Employee, error) {
	repoResp, err := e.employeeRepo.ListEmployees(page, pageSize)
	if err != nil {
		return nil, err
	}
	if len(repoResp) == 0 {
		return nil, fmt.Errorf(constants.EmployeeNotFoundError)
	}
	return repoResp, nil
}

func (e *employeeService) CreateEmployee(employee *dao.Employee) (*dao.Employee, error) {
	repoRes, err := e.employeeRepo.CreateEmployee(employee)
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (e *employeeService) GetEmployeeById(id int) (*dao.Employee, error) {
	repoRes, err := e.employeeRepo.GetEmployeeById(id)
	if err != nil {
		return nil, err
	}
	return repoRes, nil
}

func (e *employeeService) UpdateEmployee(employee *dao.Employee) error {
	err := e.employeeRepo.UpdateEmployee(employee)
	if err != nil {
		return err
	}
	return nil
}

func (e *employeeService) DeleteEmployee(id int) error {
	repoReq := &dao.Employee{
		Id: id,
	}
	return e.employeeRepo.DeleteEmployee(repoReq)
}
