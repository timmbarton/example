package cache

import (
	"github.com/timmbarton/layout/components/redisconn"
)

type cache struct {
	rd *redisconn.Conn
}

func NewCache(rd *redisconn.Conn) Cache {
	return &cache{
		rd: rd,
	}
}
