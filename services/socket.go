package services

import (
	"github.com/omidfth/godp/models"
	"net"
)

type SocketService struct {
	sockets []models.Socket
}

func NewSocketService() *SocketService {
	var sockets []models.Socket
	return &SocketService{
		sockets: sockets,
	}
}

func (s *SocketService) Get(socketId string) (*models.Socket, int) {
	for i, so := range s.sockets {
		if so.SocketID == socketId {
			return &so, i
		}
	}
	return nil, -1
}

func (s *SocketService) GetWithAddress(addr net.Addr) (*models.Socket, int) {
	for i, socket := range s.sockets {
		if socket.Address.String() == addr.String() {
			return &socket, i
		}
	}
	return nil, -1
}

func (s *SocketService) RemoveAddress(address net.Addr) *SocketService {
	socketIndex := s.GetSocketIndexByAddress(address)
	if socketIndex < 0 {
		return s
	}
	s.Remove(s.sockets[socketIndex])
	return s
}

func (s *SocketService) GetSocketIndexByAddress(address net.Addr) int {
	for i, socket := range s.sockets {
		if socket.Address.String() == address.String() {
			return i
		}
	}

	return -1
}

func (s *SocketService) Add(socket models.Socket) *SocketService {
	s.sockets = append(s.sockets, socket)
	return s
}

func (s *SocketService) Remove(socket models.Socket) *SocketService {
	if len(s.sockets) < 2 {
		s.sockets = []models.Socket{}
		return s
	}
	so, index := s.GetWithAddress(socket.Address)
	if so == nil {
		return s
	}
	s.sockets = append(s.sockets[:index], s.sockets[index+1:]...)
	return s
}
