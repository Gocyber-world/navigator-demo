package core

import (
	"os"

	"github.com/Gocyber-world/navigator-demo/initialize"
	"github.com/Gocyber-world/navigator-demo/logger"
)

func RunWindowsServer() {
	router := initialize.Routers()
	if err := router.Run("0.0.0.0:9000"); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
