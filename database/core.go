package database

import (
	"fmt"
	"log"

	"gobin/database/memcached"
	"gobin/database/postgres"
	"gobin/utils/config"
)

type Database struct {
	Memcached *memcached.Memcached
	//
	Postgres      *postgres.Postgres
	postgresClose func() error
}

func (d Database) Close() {
	log.Println("Closing database connections")
	// Postgres
	if err := d.postgresClose(); err != nil {
		log.Println("Postgres:", err)
	}
	//
}

func New(conf *config.Config) (*Database, error) {

	mc, err := memcached.New(conf.Memcached)
	if err != nil {
		return nil, fmt.Errorf("memcached: %w", err)
	}

	pg, Close, err := postgres.New(conf.Postgres)
	if err != nil {
		return nil, fmt.Errorf("posgres: %w", err)
	}

	return &Database{
		Memcached:     mc,
		Postgres:      pg,
		postgresClose: Close,
	}, nil

}
