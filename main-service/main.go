package main

import (
	"os"

	"example.com/main-service/database"
	"example.com/main-service/service"
	"github.com/alexcesaro/log/stdlog"
)

func main() {
	logger := stdlog.GetFromFlags()
	logger.Info("Exampl Service started!")

	fileData, err := os.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	database.SetParameters(fileData)
	service.StartRestService()
}
