package main

import (
	rtspService "github.com/gyeonghokim/rtsp2socket.io/rtspproxy/application"
	log "github.com/sirupsen/logrus"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	log.SetFormatter(&log.JSONFormatter{})

	// Output to stdout instead of the default stderr
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
}

func main() {
	log.WithFields(log.Fields{
		"module":   "main",
		"function": "main",
	}).Info("Starting application")
	go rtspService.RTSPServer()
	go SocketIOServer()
}
