package mysql

import (
	"github.com/kimxuanhong/go-database/core"
	"gorm.io/driver/mysql"
)

// NewDatabase initializes and returns a new Client instance.
//
// Example:
//
//	cfg := &Config{}
//	pg, err := NewDatabase(cfg)
//	if err != nil {
//	    log.Fatal(err)
//	}
//	defer pg.Close()
func NewDatabase(cfg *core.Config) (core.Database, error) {
	return core.NewDatabase(mysql.Open(cfg.GetDSN()), cfg)
}
