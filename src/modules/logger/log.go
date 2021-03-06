package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func Log(sender, message string) {
	display(sender, message)
	writeFile(sender, message)
}

func display(sender, message string) {
	log.Printf("| %s | %s\n", sender, message)
}

func writeFile(sender, message string) {
	globalFile, err := os.OpenFile(os.Getenv("LOGS"), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		Log("Logger", fmt.Sprintf("Unable to create file: %s - %s", os.Getenv("LOGS"), err.Error()))
	}
	defer globalFile.Close()
	timeString := strings.Split(fmt.Sprintf("%s", time.Now()), " ")

	if err != nil {
		Log("Logger", fmt.Sprintf("Unable to parse time: %s", err.Error()))
	}
	_, err = globalFile.WriteString(fmt.Sprintf("%v | %s %s | %s\n", timeString[0], timeString[1], sender, message))
	if err != nil {
		Log("Logger", fmt.Sprintf("Unable to write logs to global file: %s", err.Error()))
	}
}
