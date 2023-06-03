package controllername

import "github.com/alexcesaro/log/stdlog"

func MainService() {
	logger := stdlog.GetFromFlags()
	logger.Info("Main service started")
}
