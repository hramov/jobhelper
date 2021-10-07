package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	model "github.com/hramov/jobhelper/src/modules/database/model"
	"github.com/hramov/jobhelper/src/modules/logger"
)

type Gorm struct {
	db *gorm.DB
}

func (g *Gorm) Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Log("Database", err.Error())
	}
	g.db = db
	logger.Log("Database", "Successfully connected")
}

func (g *Gorm) GetConnection() *gorm.DB {
	return g.db
}

func (g *Gorm) Migrate() error {
	err := g.db.AutoMigrate(&model.Device{})
	err = g.db.AutoMigrate(&model.User{})
	err = g.db.AutoMigrate(&model.DeviceChange{})
	return err
}
