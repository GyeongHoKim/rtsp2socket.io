package repository

type RTSPStreamRepository interface {
	GetRTSPPort() string
	StreamExists(streamID string) bool
	AddClient(streamID string, client ClientST)
}
