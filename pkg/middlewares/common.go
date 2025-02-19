package middlewares

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func NoRoute(c *gin.Context) {
	errMsg := c.Request.URL.Path + " endpoint not found"
	log.Error(errMsg)
	c.AbortWithStatusJSON(http.StatusNotFound, errMsg) //Can pass user ID or request ID here for tracking purpose
}

func NoMethod(c *gin.Context) {
	errMsg := "method " + c.Request.Method + " not allowed for " + c.Request.URL.Path
	log.Error(errMsg)
	c.AbortWithStatusJSON(http.StatusNotFound, errMsg) //Can pass user ID or request ID here for tracking purpose
}
