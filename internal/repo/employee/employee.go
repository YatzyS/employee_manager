package employee

import (
	"fmt"
	"github.com/employee_manager/internal/common/constants"
	"github.com/employee_manager/internal/dao"
)

func (e *employeeRepo) CreateEmployee(data *dao.Employee) (*dao.Employee, error) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.idCounter++
	data.Id = e.idCounter
	e.inMemoryStore[e.idCounter] = data
	return data, nil
}

func (e *employeeRepo) GetEmployeeById(id int) (*dao.Employee, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	emp, ok := e.inMemoryStore[id]
	if !ok {
		return nil, fmt.Errorf(constants.EmployeeNotFoundError)
	}
	return emp, nil
}

func (e *employeeRepo) UpdateEmployee(data *dao.Employee) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.inMemoryStore[data.Id]; !ok {
		return fmt.Errorf(constants.EmployeeNotFoundError)
	}
	if data.Position != "" {
		e.inMemoryStore[data.Id].Position = data.Position
	}
	if data.Name != "" {
		e.inMemoryStore[data.Id].Name = data.Name
	}
	if data.Salary != 0 {
		e.inMemoryStore[data.Id].Salary = data.Salary
	}
	return nil
}

func (e *employeeRepo) DeleteEmployee(data *dao.Employee) error {
	e.mu.Lock()
	defer e.mu.Unlock()
	if _, ok := e.inMemoryStore[data.Id]; !ok {
		return fmt.Errorf(constants.EmployeeNotFoundError)
	}
	delete(e.inMemoryStore, data.Id)
	return nil
}

func (e *employeeRepo) ListEmployees(page int, pageSize int) ([]*dao.Employee, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	keys := getKeys(e.inMemoryStore)
	start, end, err := getStarAndEndIdx(page, pageSize, len(keys))
	if err != nil {
		return nil, err
	}
	result := make([]*dao.Employee, 0)
	for i := start; i < end; i++ {
		key := keys[i]
		result = append(result, e.inMemoryStore[key])
	}
	return result, nil
}
