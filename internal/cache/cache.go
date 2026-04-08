package cache

import (
	"github.com/timmbarton/example/internal/usecase"
	"github.com/timmbarton/layout/components/redisconn"
)

type cache struct {
	rd *redisconn.Conn
}

func NewCache(rd *redisconn.Conn) usecase.Cache {
	return &cache{
		rd: rd,
	}
}
