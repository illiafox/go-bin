package memcached

import (
	"github.com/bradfitz/gomemcache/memcache"
	"gobin/public"
)

func (m Memcached) NewBin(key string, data []byte) error {
	return public.NewInternal(
		m.client.Set(&memcache.Item{
			Key:        key,
			Value:      data,
			Flags:      0,
			Expiration: expire,
		}),
	)
}
