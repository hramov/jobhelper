package server

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/hramov/jobhelper/src/modules/server/middleware"
	"github.com/hramov/jobhelper/src/modules/server/routes"
)

type Gin struct{}

func (g *Gin) Start() {

	router := gin.Default()
	router.Static("/uploads", "uploads")
	router.Use(middleware.CORSMiddleware())

	routes.Register(router)

	err := router.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
	if err != nil {
		log.Fatal(err)
	}

}
