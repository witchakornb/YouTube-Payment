package database

import (
	"fmt"
	"sync"

	"github.com/witchakornb/YouTube-Payment/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDB struct {
	db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDB
)

func NewPostgresDatabase(config *config.Database) Database {
	once.Do(func() {
		dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=%s TimeZone=%s",
			config.Host,
			config.Port,
			config.User,
			config.Password,
			config.DatabaseName,
			config.SslMode,
			config.Timezone,
		)
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}
		dbInstance = &postgresDB{db: db}
	})
	return dbInstance
}

func (p *postgresDB) GetDB() *gorm.DB {
	return dbInstance.db
}
