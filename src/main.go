package main

import (
	"log"
	"os"

	"github.com/hramov/jobhelper/src/modules/database"
	"github.com/hramov/jobhelper/src/modules/files"
	grpc_server "github.com/hramov/jobhelper/src/modules/grpc"
	"github.com/hramov/jobhelper/src/modules/ioc"
	"github.com/hramov/jobhelper/src/modules/server"
	"github.com/hramov/jobhelper/src/modules/telegram"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()                    // Load the .env file
	files.CheckFile(os.Getenv("LOGS")) // Create log file if not exists

	orm := &database.Gorm{} // Init GORM instance
	orm.Connect()           // Connect to DB
	orm.Migrate()           // Automigrating models

	if err := ioc.Register(orm.GetConnection()); err != nil { // Register DB connection instance in IoC container
		log.Fatal("Cannot use IoC container!") // Exit from the app if IoC throws error
	}

	bot := telegram.TGBot{Token: os.Getenv("TOKEN"), Admin: "therealhramov"} // Init Telegram Bot instance with token
	bot.Create()                                                             // Create bot
	go bot.HandleQuery(bot.Update)                                           // Goroutine that handles telegram bot queries

	grpc := grpc_server.Server{}
	go grpc.Start()

	app := server.Gin{} // Init server instance
	app.Start()         // Start server to handle connections

}
