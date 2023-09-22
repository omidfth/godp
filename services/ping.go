package services

import (
	"sync"
	"time"
)

type PingService struct {
	pingMap sync.Map //map[string]time.Time
}

func NewPingService() *PingService {
	pingMap := sync.Map{}
	return &PingService{
		pingMap: pingMap,
	}
}

func (p *PingService) Remove(socketId string) *PingService {
	_, ok := p.pingMap.Load(socketId)
	if ok {
		p.pingMap.Delete(socketId)
	}
	return p
}

func (p *PingService) IsDisconnect(socketId string) bool {
	val, ok := p.pingMap.Load(socketId)
	if !ok {
		return true
	}
	if time.Now().Sub(val.(time.Time)).Seconds() > 5 {
		return true
	}
	return false
}

func (p *PingService) Ping(socketId string) *PingService {
	p.pingMap.Store(socketId, time.Now())
	return p
}
