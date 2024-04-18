package database

import (
	"time"

	"gorm.io/gorm"

	// PostgreSQL
	_ "github.com/lib/pq"
)

type (
	// db is wrapper for Master and Slave database connection
	Replication struct {
		DriverName string
		Master     *gorm.DB
		Slave      *gorm.DB
	}

	// ConnectionOptions list of option to connect to database
	ConnectionOptions struct {
		Retry                 int
		MaxOpenConnections    int
		MaxIdleConnections    int
		ConnectionMaxLifetime time.Duration
	}
)

func (db *Replication) GetMaster() *gorm.DB {
	return db.Master
}

func (db *Replication) GetSlave() *gorm.DB {
	return db.Slave
}
