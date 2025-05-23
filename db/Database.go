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

func BuildDSN(c *Config) (gorm.Dialector, error) {
	switch c.Driver {
	case "postgres":
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s search_path=%s sslmode=%s",
			c.Host, c.Port, c.User, c.Password, c.DBName, c.Schema, c.SSLMode,
		)
		return postgres.Open(dsn), nil
	case "mysql":
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
			c.User, c.Password, c.Host, c.Port, c.DBName,
		)
		return mysql.Open(dsn), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %s", c.Driver)
	}
}

func Open(configs ...*Config) (*Database, error) {
	cfg := GetConfig(configs...)
	dialect, err := BuildDSN(cfg)
	if err != nil {
		log.Printf("failed to connect database: %v", err)
		return nil, err
	}
	db, err := gorm.Open(dialect, &gorm.Config{})
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
