package server

import (
	"github.com/employee_manager/internal/handler/rest"
	"github.com/gin-gonic/gin"
)

func (a *App) SetupRoutesAndMiddleware(router *gin.RouterGroup, restHandler rest.RestHandler) {
	router.Use(gin.Recovery())
	v1 := router.Group("/v1/")
	v1.POST("/create", restHandler.CreateEmployee)
	v1.POST("/update", restHandler.UpdateEmployee)
	v1.DELETE("/delete", restHandler.DeleteEmployee)
	v1.GET("/list", restHandler.ListEmployees)
	v1.GET("/get", restHandler.GetEmployeeById)
}
