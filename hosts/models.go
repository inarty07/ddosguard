package hosts

import (
	"net"
	"time"
)

type MsgPack struct {
	Domain    string
	IP        uint32
	TimeStamp int64
}

type MsgHumanFriendly struct {
	Domain string
	IP     net.IP
	Time   time.Time
}

func (m *MsgPack) convertIntToNetIP() net.IP {
	result := make(net.IP, 4)
	result[0] = byte(m.IP)
	result[1] = byte(m.IP >> 8)
	result[2] = byte(m.IP >> 16)
	result[3] = byte(m.IP >> 24)

	return result
}
