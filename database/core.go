package database

import (
	"fmt"
	"log"

	"gobin/database/memcached"
	"gobin/database/postgres"
	"gobin/utils/config"
)

type PoolClose func()

type Database struct {
	Memcached *memcached.Memcached
	//
	Postgres *postgres.Postgres
	//
	closePool PoolClose
}

func (d Database) Close() {
	log.Println("Closing database connections")

	d.closePool()
	//
}

func New(conf *config.Config) (*Database, error) {

	mc, err := memcached.New(conf.Memcached)
	if err != nil {
		return nil, fmt.Errorf("memcached: %w", err)
	}

	pg, closePG, err := postgres.New(conf.Postgres)
	if err != nil {
		return nil, fmt.Errorf("posgres: %w", err)
	}

	return &Database{
		Memcached: mc,
		Postgres:  pg,
		closePool: closePG,
	}, nil

}
