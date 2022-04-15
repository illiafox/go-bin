package postgres

import (
	"database/sql"
	"errors"
	"fmt"
	"gobin/public"
)

func (p Postgres) GetBin(key string) ([]byte, error) {
	var data []byte

	err := p.conn.QueryRow("SELECT data FROM bins WHERE key = $1", key).Scan(&data)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, public.NewInternal(
			fmt.Errorf("select and scan data: %w", err),
		)
	}

	return data, nil
}
