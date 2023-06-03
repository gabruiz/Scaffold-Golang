package service

import (
	"example.com/main-service/endpoint"
	"github.com/alexcesaro/log/stdlog"
	"github.com/gin-gonic/gin"
)

func StartRestService() {
	logger := stdlog.GetFromFlags()
	logger.Info("Rest Service started!")

	router := gin.Default()
	endpoint.InitExampleRestService(router)

	router.Run("localhost:8083")
}
