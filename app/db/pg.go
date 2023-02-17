package db

import (
	"log"
	"sync"
	"time"

	"github.com/go-pg/pg/v10"
)

var (
	instance       *pg.DB
	once           sync.Once
	MAX_DB_RETRIES = 3
)

func NewPgDB(pgOptions *pg.Options) *pg.DB {
	once.Do(func() {
		instance = Connect(pgOptions)
	})

	return instance
}

func Connect(pgOptions *pg.Options) *pg.DB {
	log.Println("Connecting to postgres database...")
	db := pg.Connect(pgOptions)

	var (
		retires, n int
		err        error
	)

	for retires = 0; retires < MAX_DB_RETRIES; retires++ {
		if _, err = db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
			log.Fatalln(err)
			time.Sleep(2 * time.Second)
		} else {
			break
		}
	}

	if retires == MAX_DB_RETRIES {
		log.Panicf("Postgres connection error %+v\n", err)
	}

	log.Println("Connection to postgres verified and successfully connected...")
	return db
}
