package storage

import (
	"fmt"
	"log"
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	_ "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

// Driver of storage
type Driver string

// Drivers
const (
	MySQL Driver = "MYSQL"
	Postgres Driver = "POSTGRES"
)

// New create the connection with db
func New(d Driver) {
	switch d {
	case MySQL:
		newMySQLDB()
	case Postgres:
		newPostgresDB()
	}
}

func newPostgresDB() {
	once.Do(func(){
		var err error
		dsn := "postgres://postgres:admin@localhost:5432/godb?sslmode=disable"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Connected to postgress")
	})
}

func newMySQLDB() {
	once.Do(func(){
		var err error
		dsn := "root:@tcp(localhost:3306)/godb?parseTime=true"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't open db: %v", err)
		}

		fmt.Println("Connected to mySQL")
	})
}

// DB return a unique instance of db
func DB() *gorm.DB {
	return db
}