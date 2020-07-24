package model

import (
	"fmt"
	"time"
)

type IP struct {
	Ip         string
	Type       string
	Location   string
	Speed      float64
	LastVerify time.Time
}

func (p IP) String() string {
	return fmt.Sprintf("ip: [%v], type: %v, location: %v, speed: %v, lastCheck: %v\n", p.Ip, p.Type, p.Location, p.Speed, p.LastVerify)
}
