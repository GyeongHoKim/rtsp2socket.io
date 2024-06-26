package application

import "net"

type RTSPStreamService interface {
	RTSPServer()
	RTSPServerClientHandle(conn net.Conn)
	RTSPServerClientPlay(uuid string, channel string, conn net.Conn)
}
