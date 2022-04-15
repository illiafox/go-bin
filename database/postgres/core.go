package postgres

import (
	"database/sql"
	"fmt"

	// postgres
	_ "github.com/lib/pq"
	//
	"gobin/utils/config"
)

type Postgres struct {
	conn *sql.DB
}

func New(conf config.Postgres) (*Postgres, func() error, error) {
	conn, err := sql.Open(
		"postgres",
		fmt.Sprintf("postgres://%s:%s@%v:%v/%v?sslmode=disable",
			conf.User,
			conf.Pass,
			conf.IP,
			conf.Port,
			conf.DbName,
		),
	)

	if err != nil {
		return nil, nil, fmt.Errorf("opening connection: %w", err)
	}

	err = conn.Ping()
	if err != nil {
		return nil, nil, fmt.Errorf("ping: %w", err)
	}

	return &Postgres{conn}, conn.Close, nil
}
