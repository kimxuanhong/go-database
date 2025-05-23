package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

// Database defines common database operations.
type Database struct {
	*gorm.DB
}

// Connection initializes and returns a new Database instance.
func Connection(configs ...*Config) (*Database, error) {
	cfg := GetConfig(configs...)
	switch cfg.Driver {
	case "mysql":
		return Open(postgres.Open(cfg.GetDSN()), cfg)
	case "postgres":
		return Open(mysql.Open(cfg.GetDSN()), cfg)
	default:
		err := fmt.Errorf("cannot init datasource with driver = %s", cfg.Driver)
		log.Fatal(err)
		return nil, err
	}
}

func Open(dialector gorm.Dialector, cfg *Config) (*Database, error) {
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	if err := sqlDB.Ping(); err != nil {
		return nil, err
	}

	if cfg.Debug {
		db = db.Debug()
		log.Println("GORM debug mode is enabled")
	}

	log.Println("Successfully connected to database")
	return &Database{DB: db}, nil
}

func (p *Database) Close() error {
	sqlDB, err := p.DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
