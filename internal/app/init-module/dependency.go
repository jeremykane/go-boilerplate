package init_module

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/jeremykane/go-boilerplate/internal/config"
	"github.com/jeremykane/go-boilerplate/pkg/database"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitializeDB(dbConfig map[string]*config.Database) map[string]*database.Replication {
	const funcName = "InitializeDB"

	replicaCollection := make(map[string]*database.Replication)
	ctx := context.Background()
	for dbName, conf := range dbConfig {
		if conf == nil {
			log.Printf("[%s] empty config for [%s]", funcName, dbName)
			continue
		}

		var (
			replica *database.Replication
			err     error
		)

		replica = &database.Replication{}

		replica.Master, err = Connect(ctx, conf.Master.Url, 1)
		if err != nil {
			log.Fatal("failed to connect to master DB", err)
		}

		// initialize Slave
		if conf.Slave != nil {
			replica.Slave, err = Connect(ctx, conf.Slave.Url, 1)
			if err != nil {
				log.Fatal("failed to connect to slave DB", err)
			}
		}
		replicaCollection[dbName] = replica
	}

	fmt.Println("all database initiated")
	return replicaCollection
}

func connectWithRetry(ctx context.Context, dataSourceName string, retry int) (*gorm.DB, error) {
	var (
		db  *gorm.DB
		err error
	)

	for t := 0; t <= retry; t++ {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN: dataSourceName,
			//DSN:                  "user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai",
			PreferSimpleProtocol: true, // disables implicit prepared statement usage
		}), &gorm.Config{})
		if err != nil {
			log.Fatal(err.Error())
			time.Sleep(time.Second * 3)
		} else {
			break
		}
	}

	return db, err
}

// Connect to a database
func Connect(ctx context.Context, dataSourceName string, retry int) (*gorm.DB, error) {
	db, err := connectWithRetry(ctx, dataSourceName, retry)
	if err != nil {
		log.Fatal("failed to connect DB ", err)
	}
	return db, nil
}
