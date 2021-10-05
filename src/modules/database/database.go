package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	device_db "github.com/hramov/jobhelper/src/modules/database/device"
)

type Gorm struct {
	db *gorm.DB
}

func (g *Gorm) Connect() {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"), os.Getenv("DB_SSL_MODE"))
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	g.db = db
}

func (g *Gorm) GetConnection() *gorm.DB {
	return g.db
}

func (g *Gorm) Migrate() error {
	err := g.db.AutoMigrate(&device_db.Device{})
	return err
}
