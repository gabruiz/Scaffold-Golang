package main

import (
	"io/ioutil"

	"example.com/main-service/database"
	"github.com/alexcesaro/log/stdlog"
)

func main() {
	logger := stdlog.GetFromFlags()
	logger.Info("Exampl Service started!")

	fileData, err := ioutil.ReadFile("./config/config.json")
	if err != nil {
		panic(err)
	}

	database.SetParameters(fileData)
}
