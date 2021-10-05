package controller

import (
	"strconv"

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

func (dc *DeviceController) FindByID(c *gin.Context) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	device_id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{"data": nil, "error": err})
		return
	}

	device, err := deviceEntity.ShowDeviceByID(uint(device_id))
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{"data": nil, "error": err})
		return
	}

	c.JSON(200, gin.H{"data": device, "error": err})
}

func (dc *DeviceController) FindExpires(c *gin.Context) {
	var deviceEntity device_core.DeviceEntityPort
	container.NamedResolve(&deviceEntity, "DeviceEntity")

	devices, err := deviceEntity.ShowAllDevices()
	if err != nil {
		c.AbortWithStatusJSON(200, gin.H{"data": nil, "error": err})
	}

	c.JSON(200, gin.H{"data": devices, "error": err})
}
