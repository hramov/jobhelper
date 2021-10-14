package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type MainController struct{}

func (mc *MainController) HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "OK"})
}
