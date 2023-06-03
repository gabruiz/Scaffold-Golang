package controller

import "github.com/alexcesaro/log/stdlog"

func ExampleService() {
	logger := stdlog.GetFromFlags()
	logger.Info("Example service started")

}
