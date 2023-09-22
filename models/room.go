package models

type Room struct {
	RoomID  string   `json:"r"`
	Sockets []Socket `json:"s"`
}
