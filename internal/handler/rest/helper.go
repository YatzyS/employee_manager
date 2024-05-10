package rest

import (
	"fmt"
	"github.com/employee_manager/internal/common/constants"
	"github.com/employee_manager/internal/dao"
	"github.com/gin-gonic/gin"
	"strconv"
)

func extractId(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil || id < 1 {
		if id < 1 {
			err = fmt.Errorf(constants.InvalidIdError)
		}
		return 0, err
	}
	return id, nil
}

func extractEmployee(c *gin.Context) (*dao.Employee, error) {
	var employee dao.Employee
	if err := c.BindJSON(&employee); err != nil {
		return nil, err
	}
	return &employee, nil
}
