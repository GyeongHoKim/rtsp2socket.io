package domain

import (
	"github.com/deepch/vdk/av"
	"net"
	"time"
)

type RTSPServerST struct {
	Debug    bool   `json:"debug"`
	RTSPPort string `json:"rtsp_port"`
}

type RTSPStreamST struct {
	StreamID string `json:"stream_id"`
	Url      string `json:"url"`
	Debug    bool   `json:"debug"`
	Status   string `json:"status"`
	Audio    bool   `json:"audio"`
	runLock  bool
	codecs   []av.CodecData
	sdp      []byte
	signals  chan int
	clients  map[string]ClientST
	ack      time.Time
}

type ClientST struct {
	mode              int
	signals           chan int
	outgoingAVPacket  chan *av.Packet
	outgoingRTPPacket chan *[]byte
	socket            net.Conn
}
