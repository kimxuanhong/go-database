package postgres

import (
	"github.com/kimxuanhong/go-database/core"
	"gorm.io/driver/postgres"
)

// NewDatabase initializes and returns a new Client instance.
//
// Example:
//
//	cfg := &Config{}
//	pg, err := NewPostgres(cfg)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer pg.Close()
func NewDatabase(cfg *core.Config) (core.Database, error) {
	return core.NewDatabase(postgres.Open(cfg.GetDSN()), cfg)
}
