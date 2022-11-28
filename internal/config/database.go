package config

import (
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/octaviomuller/kendamais-server/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var once sync.Once
var mu sync.Mutex

func GetDB() *gorm.DB {
	return db
}

func RunMigrationsDB() (err error) {
	err = db.AutoMigrate(model.User{}, model.Bidding{})
	if err != nil {
		return err
	}

	return
}

func ConnectDB(connectionString string) *gorm.DB {
	mu.Lock()
	defer mu.Unlock()

	once.Do(func() {
		var err error
		attempts := 3

		// Attempts to stablish connection
		for attempts > 0 {
			db, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
			if err != nil {
				fmt.Println("Error while connecting to database. Trying again in 5 seconds")
			} else {
				break
			}

			<-time.After(5 * time.Second)

			attempts--
		}

		// Throwing fatal if some error occured
		if err != nil {
			log.Fatal("Error while connecting to database")
		}

		// Applying migrations
		err = RunMigrationsDB()
		if err != nil {
			log.Fatal("Error while migrate")
		}
	})

	return db
}
