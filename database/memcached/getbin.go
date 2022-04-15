package memcached

import (
	"encoding/json"
	"errors"
	"fmt"
	"gobin/database/model"
	"gobin/public"

	"github.com/bradfitz/gomemcache/memcache"
)

// GetBin gets value by key and updates its expiration time
func (m Memcached) GetBin(key string) (*model.Bin, error) {

	item, err := m.client.Get(key)

	if err != nil {
		if errors.Is(err, memcache.ErrCacheMiss) {
			return nil, nil
		}

		return nil, public.NewInternal(
			fmt.Errorf("get cache: %w", err),
		)
	}

	if item.Expiration < update {
		err = m.client.Touch(key, expire)
		if err != nil {
			return nil, public.NewInternal(
				fmt.Errorf("touch (update expiration) cache: %w", err),
			)
		}
	}

	var bin model.Bin

	err = json.Unmarshal(item.Value, &bin)
	if err != nil {
		return nil, public.NewInternal(
			fmt.Errorf("unmarshalling cache: %w", err),
		)
	}

	return &bin, nil
}
