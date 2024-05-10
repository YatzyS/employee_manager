package employee

import (
	"fmt"
	"github.com/employee_manager/internal/dao"
	"testing"
)

func TestEmployeeRepo(t *testing.T) {
	mockRepo := NewEmployeeRepo()
	employee := &dao.Employee{
		Name:     "John Doe",
		Position: "Manager",
		Salary:   50000,
	}
	createdEmployee, err := mockRepo.CreateEmployee(employee)
	if err != nil {
		t.Errorf("CreateEmployee failed: %v", err)
	}
	if createdEmployee == nil {
		t.Error("CreateEmployee returned nil employee")
	}

	// Test GetEmployeeById
	retrievedEmployee, err := mockRepo.GetEmployeeById(createdEmployee.Id)
	if err != nil {
		t.Errorf("GetEmployeeById failed: %v", err)
	}
	if retrievedEmployee == nil {
		t.Error("GetEmployeeById returned nil employee")
	}

	// Test UpdateEmployee
	retrievedEmployee.Position = "Senior Manager"
	err = mockRepo.UpdateEmployee(retrievedEmployee)
	if err != nil {
		t.Errorf("UpdateEmployee failed: %v", err)
	}
	retrievedEmployee, err = mockRepo.GetEmployeeById(retrievedEmployee.Id)
	if err != nil {
		t.Errorf("GetEmployeeById failed: %v", err)
	}
	if retrievedEmployee == nil {
		t.Error("GetEmployeeById returned nil employee")
	}
	if retrievedEmployee.Position != "Senior Manager" {
		t.Error("Employee not updated")
	}

	// Test DeleteEmployee
	err = mockRepo.DeleteEmployee(retrievedEmployee)
	if err != nil {
		t.Errorf("DeleteEmployee failed: %v", err)
	}
	retrievedEmployee, err = mockRepo.GetEmployeeById(retrievedEmployee.Id)
	if retrievedEmployee != nil {
		t.Error("Employee not deleted")
	}

	// List employee negative case
	employees, err := mockRepo.ListEmployees(1, 5)
	if err == nil {
		t.Errorf("ListEmployees negative case failed: %v", err)
	}

	// Test GetEmployeeById Negative case
	retrievedEmployee, err = mockRepo.GetEmployeeById(createdEmployee.Id)
	if err == nil {
		t.Errorf("GetEmployeeById negative case failed: %v", err)
	}

	// Test DeleteEmployee Negative case
	err = mockRepo.DeleteEmployee(createdEmployee)
	if err == nil {
		t.Errorf("GetEmployeeById negative case failed: %v", err)
	}

	// Test UpdateEmployee negative case
	err = mockRepo.UpdateEmployee(createdEmployee)
	if err == nil {
		t.Errorf("UpdateEmployee negative case failed: %v", err)
	}

	for i := 1; i <= 10; i++ {
		mockRepo.CreateEmployee(&dao.Employee{
			Id:       i,
			Name:     fmt.Sprintf("Employee%d", i),
			Position: "Developer",
			Salary:   60000,
		})
	}
	// Test pagination for the first page with page size 5
	employees, err = mockRepo.ListEmployees(1, 5)
	if err != nil {
		t.Errorf("ListEmployees failed: %v", err)
	}
	if len(employees) != 5 {
		t.Errorf("ListEmployees returned wrong number of employees, expected 5, got %d", len(employees))
	}
}
