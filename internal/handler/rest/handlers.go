package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (r *restHandler) ListEmployees(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page < 1 {
		page = 1
	}
	pageSize, err := strconv.Atoi(c.Query("pageSize"))
	if err != nil || pageSize < 1 {
		pageSize = 10
	}
	employees, err := r.employeeService.ListEmployees(page, pageSize)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employees)
}

func (r *restHandler) CreateEmployee(c *gin.Context) {
	employee, err := extractEmployee(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	newEmployee, err := r.employeeService.CreateEmployee(employee)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, newEmployee)
}

func (r *restHandler) GetEmployeeById(c *gin.Context) {
	id, err := extractId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	employee, err := r.employeeService.GetEmployeeById(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, employee)
}

func (r *restHandler) DeleteEmployee(c *gin.Context) {
	id, err := extractId(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err = r.employeeService.DeleteEmployee(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}

func (r *restHandler) UpdateEmployee(c *gin.Context) {
	employee, err := extractEmployee(c)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	err = r.employeeService.UpdateEmployee(employee)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Message": "Success"})
}
