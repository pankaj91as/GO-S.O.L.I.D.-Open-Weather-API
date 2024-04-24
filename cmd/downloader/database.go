package main

import (
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DB interface {
	Connect() (*gorm.DB, *sql.DB, error)
	Close(db *sql.DB) error
}

type SQLConnection struct {
	host     string
	port     int
	username string
	password string
	database string
}

func MySQLConnect(host string, port int, username string, password string, database string) *SQLConnection {
	return &SQLConnection{
		host:     host,
		port:     port,
		username: username,
		password: password,
		database: database,
	}
}

func (cs *SQLConnection) Connect() (*gorm.DB, *sql.DB, error) {
	// dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cs.username, cs.password, cs.host, cs.port, cs.database)
	// db, err := sql.Open("mysql", dataSourceName)
	// if err != nil {
	// 	return nil, err
	// }
	// return db, nil

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cs.username, cs.password, cs.host, cs.port, cs.database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	dbInstance, _ := db.DB()

	return db, dbInstance, nil
}

func Close(db *sql.DB) (err error) {
	return db.Close()
}
