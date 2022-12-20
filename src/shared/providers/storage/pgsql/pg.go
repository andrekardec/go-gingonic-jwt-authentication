package pgsql

import (
	"github.com/go-pg/pg/v10"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

var (
	instance *pg.DB
	once     sync.Once
)

func NewPgDb(options *pg.Options) *pg.DB {
	once.Do(func() {
		instance = Connect(options)
	})

	return instance
}

func Connect(options *pg.Options) *pg.DB {
	log.Infoln("Connecting to Postgres database...")

	db := pg.Connect(options)

	var n int

	if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
		log.Errorln(err)
		time.Sleep(2 * time.Second)

		db = pg.Connect(options)

		if _, err := db.QueryOne(pg.Scan(&n), "SELECT 1"); err != nil {
			log.Panicf("Postgres connection error %v\n", err)
		}

		log.Infoln("Connection to Postgres verified and successfully connected...")
	}

	return db
}
