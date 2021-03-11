package hosts

import (
	"encoding/json"
	"fmt"
	"time"
)

func (c *cache) Add(msg MsgPack) {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[msg.IP]; !ok {
		msg.TimeStamp = time.Now().Unix()
		c.m[msg.IP] = msg
	}
}

func (c *cache) Del(msg MsgPack) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.m, msg.IP)
}

func (c *cache) Contains(msg MsgPack) bool {
	c.mu.Lock()
	defer c.mu.Unlock()

	if _, ok := c.m[msg.IP]; ok {
		return true
	} else {
		return false
	}
}

func (c *cache) List() []MsgPack {
	c.mu.Lock()
	defer c.mu.Unlock()

	resp := make([]MsgPack, 0, len(c.m))
	for i := range c.m {
		resp = append(resp, c.m[i])
	}

	return resp
}

func (c *cache) ListJSON() ([]byte, error) {
	data, err := json.Marshal(c.m)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal map: %w", err)
	}

	return data, err
}

func (c *cache) ListHumanFreindly() map[uint32]MsgHumanFriendly {
	c.mu.Lock()
	defer c.mu.Unlock()

	humanFriendlyMap := make(map[uint32]MsgHumanFriendly, len(c.m))
	for k, v := range c.m {
		humanFriendlyMap[k] = MsgHumanFriendly{
			Domain: v.Domain,
			IP:     v.convertIntToNetIP(),
			Time:   time.Unix(v.TimeStamp, 0),
		}
	}

	return humanFriendlyMap
}

func (c *cache) Clear() {
	c.m = make(map[uint32]MsgPack, 0)
}

func (c *cache) Len() int {
	c.mu.Lock()
	defer c.mu.Unlock()

	return len(c.m)
}
