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
	GormConn *gorm.DB
	SqlCon   *sql.DB
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

func CreateTable(ormdb *gorm.DB, tableNames []string, ns interface{}) error {
	for _, tableName := range tableNames {
		err := ormdb.Table(string(tableName)).AutoMigrate(ns)
		if err != nil {
			return err
		}
	}

	return nil
}
