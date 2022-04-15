package postgres

import "gobin/public"

func (p Postgres) NewBin(key string, data []byte) error {
	_, err := p.conn.Exec("INSERT INTO bins (key, data) VALUES ($1,$2)", key, data)

	return public.NewInternal(err)
}
