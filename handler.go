package godp

import (
	"encoding/json"
	"fmt"
	"godp/internal/debugger"
	"godp/internal/events"
	"godp/models"
	"godp/services"
	"log"
	"net"
)

var (
	packetConn    net.PacketConn
	packetError   error
	maxBufferSize int
)

const udpDebugTag = "UDP::"

func NewRouter() *Router {
	socketService := services.NewSocketService()
	roomService := services.NewRoomService()
	pingService := services.NewPingService()
	router := Router{
		eventRoutes:   make(map[uint]*Route),
		SocketService: socketService,
		RoomService:   roomService,
		PingService:   pingService,
	}
	return &router
}

func (r *Router) ListenAndServe(port uint, bufferSize int) {
	maxBufferSize = bufferSize
	packetConn, packetError = net.ListenPacket("udp", fmt.Sprintf("0.0.0.0:%d", port))
	if packetError != nil {
		log.Fatal(packetError.Error())
	}
	defer func(packetConn net.PacketConn) {
		var err = packetConn.Close()
		if err != nil {
			debugger.Debug(udpDebugTag, debugger.ERROR, "ERR: ", err.Error())
		}
	}(packetConn)
	serve(r)
}

func serve(router *Router) {
	debugger.Debug(udpDebugTag, 0, "Start UDP Server:", packetConn.LocalAddr())
	for {
		buf := make([]byte, maxBufferSize)
		n, addr, err := packetConn.ReadFrom(buf)
		receiver(router, addr, buf[:n])
		if err != nil {
			debugger.Debug(udpDebugTag, debugger.ERROR, "ERR: ", err.Error())
		}
	}
}

type Router struct {
	eventRoutes   map[uint]*Route
	SocketService *services.SocketService
	RoomService   *services.RoomService
	PingService   *services.PingService
}

type Route struct {
	Handler udpHandler
}

func (r *Router) NewRoute(eventId uint, f func(ctx *Context)) *Route {
	return r.addHandler(eventId, udpFunction(f))
}

func (r *Router) addHandler(eventId uint, handler udpHandler) *Route {
	route := Route{
		Handler: handler,
	}
	r.eventRoutes[eventId] = &route
	return &route
}

type udpHandler interface {
	ServeUDPFunc(*Context)
}

type udpFunction func(*Context)

func (f udpFunction) ServeUDPFunc(c *Context) {
	f(c)
}

func receiver(router *Router, addr net.Addr, buf []byte) {
	var packet models.Packet
	socket, _ := router.SocketService.GetWithAddress(addr)
	err := json.Unmarshal(buf, &packet)
	if err != nil {
		debugger.Debug(udpDebugTag, debugger.ERROR, "UDP ERR", err.Error())
	} else {
		if packet.EventID == events.PING {
			debugger.Debug(udpDebugTag, debugger.DEBUG, "RECEIVE: ", packet.EventID, " DATA: ", string(buf))
		} else {
			debugger.Debug(udpDebugTag, debugger.TRACE, "RECEIVE: ", packet.EventID, " DATA: ", string(buf))
		}
		context := Context{
			Socket:  socket,
			Address: addr,
			Packet:  packet,
			router:  router,
		}
		router.eventRoutes[packet.EventID].Handler.ServeUDPFunc(&context)
	}
}
