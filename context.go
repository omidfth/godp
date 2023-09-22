package godp

import (
	"audio_stream/events"
	"encoding/json"
	"godp/internal/debugger"
	"godp/models"
	"godp/services"
	"net"
)

type Context struct {
	Socket  *models.Socket
	Address net.Addr
	Packet  models.Packet
	router  *Router
}

func (c *Context) GetSocketService() *services.SocketService {
	return c.router.SocketService
}

func (c *Context) GetRoomService() *services.RoomService {
	return c.router.RoomService
}

func (c *Context) GetPingService() *services.PingService {
	return c.router.PingService
}

func (c *Context) Emit(addr net.Addr, buf []byte) {
	var packet models.Packet
	err := json.Unmarshal(buf, &packet)
	if err == nil && packet.EventID != events.PING {
		debugger.Debug(udpDebugTag, 0, "EMIT: ", string(buf))
	}
	packetConn.WriteTo(buf, addr)
}

func (c *Context) EmitTo(roomId string, buf []byte) {
	room, _ := c.router.RoomService.Get(roomId)
	if room != nil {
		for _, socket := range room.Sockets {
			c.Emit(socket.Address, buf)
		}
	}
}
