package application

import "net"

func RTSPServer() {
}

func RTSPServerClientHandle(conn net.Conn) {
}

func RTSPServerClientPlay(uuid string, channel string, conn net.Conn) {
}

func parsecSEQ(buf []byte) int {}

func parseStage(buf []byte) (string, error) {}

func pasrseStreamChannel(buf []byte) (string, string, string, error) {}
