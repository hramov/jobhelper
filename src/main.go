package main

import (
	"log"

	"github.com/hramov/jobhelper/src/modules/database"
	"github.com/hramov/jobhelper/src/modules/ioc"
	"github.com/hramov/jobhelper/src/modules/server"
	"github.com/joho/godotenv"
)

func main() {

	godotenv.Load()

	orm := &database.Gorm{}
	orm.Connect()
	orm.Migrate()

	if err := ioc.Register(orm.GetConnection()); err != nil {
		log.Fatal("Cannot use IoC container!")
	}

	app := server.Gin{}
	app.Start()
}
