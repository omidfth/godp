package godp

import "github.com/omidfth/godp/models"

func MakePacket(socketId string, roomId string, eventID uint, data interface{}) models.Packet {
	return models.Packet{
		SocketID: socketId,
		EventID:  eventID,
		Data:     data,
		RoomID:   roomId,
	}
}

const (
	PING       = 0
	CONNECT    = 1
	DISCONNECT = 2
	REGISTER   = 3
	JOIN       = 4
	LEAVE      = 5
	STREAM     = 6
)
