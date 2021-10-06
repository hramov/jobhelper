package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

var (
	globalFileName = "/srv/jobhelper/logs.txt"
)

func Log(sender, message string) {
	display(sender, message)
	writeFile(sender, message)
}

func display(sender, message string) {
	log.Printf("| %s | %s\n", sender, message)
}

func writeFile(sender, message string) {
	// localFile, err := os.OpenFile(localFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	globalFile, err := os.OpenFile(globalFileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0777)
	if err != nil {
		Log("Logger", fmt.Sprintf("Unable to create file: %s", err.Error()))
	}
	// defer localFile.Close()
	defer globalFile.Close()

	// _, err = localFile.WriteString(fmt.Sprintf("%v | %s | %s\n", time.Now(), sender, message))
	// if err != nil {
	// 	Log("Logger", fmt.Sprintf("Unable to write logs to local file: %s", err.Error()))
	// }

	timeString := strings.Split(fmt.Sprintf("%s", time.Now()), " ")

	if err != nil {
		Log("Logger", fmt.Sprintf("Unable to parse time: %s", err.Error()))
	}

	_, err = globalFile.WriteString(fmt.Sprintf("%v | %s %s | %s\n", timeString[0], timeString[1], sender, message))
	if err != nil {
		Log("Logger", fmt.Sprintf("Unable to write logs to global file: %s", err.Error()))
	}
}
