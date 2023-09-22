package models

type Packet struct {
	SocketID string      `json:"s"`
	EventID  uint        `json:"e"`
	Data     interface{} `json:"d"`
	RoomID   string      `json:"r"`
}
