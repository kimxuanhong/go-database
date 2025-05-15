package database

import (
	"fmt"
	"github.com/kimxuanhong/go-database/core"
	"github.com/kimxuanhong/go-database/mysql"
	"github.com/kimxuanhong/go-database/postgres"
	"log"
)

// NewDatabase initializes and returns a new Client instance.
func NewDatabase(configs ...*core.Config) (core.Database, error) {
	cfg := core.GetConfig(configs...)
	switch cfg.Driver {
	case "mysql":
		return mysql.NewDatabase(cfg)
	case "postgres":
		return postgres.NewDatabase(cfg)
	default:
		err := fmt.Errorf("cannot init datasource with driver = %s", cfg.Driver)
		log.Fatal(err)
		return nil, err
	}
}
