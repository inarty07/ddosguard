package hosts

import (
	"sync"
	"time"
)

type Hosts interface {
	Add(msg MsgPack)                                // добавление данных (аргумент в формате msgpack)
	Del(msg MsgPack)                                // удаление данных (аргумент в формате msgpack)
	Len() int                                       // кол-во объектов
	List() []MsgPack                                // список всех данных (результат в формате msgpack)
	Clear()                                         // очистка всех данных
	Contains(msg MsgPack) bool                      // проверка существования записи (аргумент в формате msgpack, результат - bool)
	ListJSON() ([]byte, error)                      // список всех данных (результат в формате json)
	ListHumanFreindly() map[uint32]MsgHumanFriendly // список всех данных (результат в формате списка из объектов типа собственной структуры)
}

type cache struct {
	m   map[uint32]MsgPack
	mu  sync.Mutex
	ttl time.Duration
}

func New(ttl time.Duration) Hosts {
	c := &cache{
		m:   make(map[uint32]MsgPack, 0),
		ttl: ttl,
	}
	go c.checkTTL()

	return c
}

func (c *cache) checkTTL() {
	for now := range time.Tick(time.Second) {
		c.mu.Lock()
		for k, v := range c.m {
			if now.Unix()+int64(c.ttl/time.Second) > v.TimeStamp {
				delete(c.m, k)
			}
		}
		c.mu.Unlock()
	}
}
