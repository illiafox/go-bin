package memcached

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
	"gobin/utils/config"
)

const (
	expire = 1800 // 30 minutes
	update = 900  // touch (update expiration) if bin will expire after 15 minutes
)

type Memcached struct {
	client *memcache.Client
}

func New(conf config.Memcached) (*Memcached, error) {
	mc := memcache.New(fmt.Sprintf("%s:%s", conf.IP, conf.Port))

	if err := mc.Ping(); err != nil {
		return nil, fmt.Errorf("memchache: ping: %w", err)
	}

	return &Memcached{mc}, nil
}
