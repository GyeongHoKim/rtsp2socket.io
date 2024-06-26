package presentation

import (
	socketio "github.com/googollee/go-socket.io"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func SocketIOServer() {
	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		log.WithFields(log.Fields{
			"module":   "SocketIOServer",
			"function": "SocketIOServer",
		}).Info("connected:", s.ID())
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		server.Remove(s.ID())
		log.WithFields(log.Fields{
			"module":   "SocketIOServer",
			"function": "SocketIOServer",
		}).Info("closed:", reason)
	})

	go server.Serve()
	defer server.Close()

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":4000", nil))
}
