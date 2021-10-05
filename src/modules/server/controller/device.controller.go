package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golobby/container/v3"
	device_core "github.com/hramov/jobhelper/src/core/device"
)

type DeviceController struct{}

func (dc *DeviceController) FindAll(c *gin.Context) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	devices, err := deviceEntity.ShowAllDevices()
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{"data": nil, "error": err})
	}

	c.JSON(200, gin.H{"data": devices, "error": err})
}
