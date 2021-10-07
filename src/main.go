package main

import (
	"log"
	"os"

	"github.com/hramov/jobhelper/src/modules/database"
	"github.com/hramov/jobhelper/src/modules/files"
	"github.com/hramov/jobhelper/src/modules/ioc"
	"github.com/hramov/jobhelper/src/modules/server"
	"github.com/hramov/jobhelper/src/modules/telegram"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	files.CheckFile(os.Getenv("LOGS"))

	orm := &database.Gorm{}
	orm.Connect()
	orm.Migrate()

	if err := ioc.Register(orm.GetConnection()); err != nil {
		log.Fatal("Cannot use IoC container!")
	}

	bot := telegram.TGBot{Token: os.Getenv("TOKEN")}
	bot.Create()

	go bot.HandleQuery(bot.Update)

	app := server.Gin{}
	app.Start()
}
