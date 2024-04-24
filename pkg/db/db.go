package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/op/go-logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Log = logging.MustGetLogger("API")

type DB interface {
	Connect() (*gorm.DB, *sql.DB, error)
	Close(db *sql.DB) error
	InitDBConnection() *SQLConnection
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

func InitDBConnection() *SQLConnection {
	// Initialize MySQL connector
	port, _ := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	connector := MySQLConnect(os.Getenv("MYSQL_HOST"), port, os.Getenv("MYSQL_USERNAME"), os.Getenv("MYSQL_PASSWORD"), "open_weather")

	// Connect to MySQL
	ormdb, dbConn, err := connector.Connect()
	if err != nil {
		Log.Fatal("Error connecting to MySQL:", err)
	}

	return &SQLConnection{
		GormConn: ormdb,
		SqlCon:   dbConn,
	}
}
