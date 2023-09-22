package godp

import (
	"godp/udp"
)

func NewUDPRouter() *udp.Router {
	return udp.NewRouter()
}
