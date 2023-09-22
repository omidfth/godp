package services

import (
	"godp/models"
	"net"
)

type RoomService struct {
	rooms []models.Room
}

func NewRoomService() *RoomService {
	var rooms []models.Room
	return &RoomService{
		rooms: rooms,
	}
}
func (r *RoomService) Rooms() []models.Room {
	return r.rooms
}

func (r *RoomService) AddSocketToRoom(roomId string, socket models.Socket) *models.Room {
	room, index := r.Get(roomId)
	if room == nil {
		return nil
	}
	r.rooms[index].Sockets = append(r.rooms[index].Sockets, socket)
	return &r.rooms[index]
}

func (r *RoomService) Add(room models.Room) *RoomService {
	r.rooms = append(r.rooms, room)
	return r
}

func (r *RoomService) Get(roomId string) (*models.Room, int) {
	for i, room := range r.rooms {
		if room.RoomID == roomId {
			return &room, i
		}
	}
	return nil, -1
}

func (r *RoomService) GetRoomWithAddress(addr net.Addr) *models.Room {
	for _, room := range r.rooms {
		for _, socket := range room.Sockets {
			if socket.Address.String() == addr.String() {
				return &room
			}
		}
	}

	return nil
}
func (r *RoomService) Remove(roomId string) *RoomService {
	if len(r.rooms) < 2 {
		r.rooms = []models.Room{}
		return r
	}
	room, index := r.Get(roomId)
	if room != nil {
		r.rooms = append(r.rooms[:index], r.rooms[index+1:]...)
		return r
	}
	return r
}

func (r *RoomService) RemoveAddress(address net.Addr) *RoomService {
	roomIndex, socketIndex := r.GetRoomAndSocketIndexByAddress(address)
	if roomIndex < 0 {
		return r
	}

	if len(r.rooms[roomIndex].Sockets) < 2 {
		r.Remove(r.rooms[roomIndex].RoomID)
		return r
	}
	r.rooms[roomIndex].Sockets = append(r.rooms[roomIndex].Sockets[:socketIndex], r.rooms[roomIndex].Sockets[socketIndex+1:]...)
	return r
}

func (r *RoomService) RemoveSocket(socket models.Socket) *RoomService {
	roomIndex, socketIndex := r.GetRoomAndSocketIndex(socket)
	if roomIndex < 0 {
		return r
	}

	if len(r.rooms[roomIndex].Sockets) < 2 {
		r.Remove(r.rooms[roomIndex].RoomID)
		return r
	}
	r.rooms[roomIndex].Sockets = append(r.rooms[roomIndex].Sockets[:socketIndex], r.rooms[roomIndex].Sockets[socketIndex+1:]...)
	return r
}

func (r *RoomService) GetRoomAndSocketIndexByAddress(address net.Addr) (int, int) {
	for i, room := range r.rooms {
		for j, socket := range room.Sockets {
			if socket.Address.String() == address.String() {
				return i, j
			}
		}
	}
	return -1, -1
}

func (r *RoomService) GetRoomAndSocketIndex(socket models.Socket) (int, int) {
	for i, room := range r.rooms {
		for j, s := range room.Sockets {
			if s.SocketID == socket.SocketID {
				return i, j
			}
		}
	}
	return -1, -1
}

func (r *RoomService) GetSocketIndexInRoom(roomIndex uint, socket models.Socket) int {
	for i, s := range r.rooms[roomIndex].Sockets {
		if s.SocketID == socket.SocketID {
			return i
		}
	}
	return -1
}
