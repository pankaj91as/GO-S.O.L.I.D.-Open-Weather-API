package db

import (
	"database/sql"
	"fmt"

	"github.com/op/go-logging"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Log = logging.MustGetLogger("API")

type DB interface {
	Connect() (*gorm.DB, *sql.DB, error)
	Close(db *sql.DB) error
	InitDBConnection(host string, port int, username string, password string) *SQLConnection
	CreateTable(ormdb *gorm.DB, tableNames []string, ns interface{}) error
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

// func (cs *SQLConnection) Connect() (*gorm.DB, *sql.DB, error) {
// 	// prepare dsn connection string
// 	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cs.username, cs.password, cs.host, cs.port, cs.database)
// 	Log.Info("Trying to connect: ", dsn)

// 	// Retry if DB not connected
// 	var err error
// 	for i := 2; i < 5; i++ {
// 		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 		if err == nil {
// 			// Connection successful
// 			Log.Info("Database connected successfuly!")
// 			dbInstance, _ := db.DB()
// 			return db, dbInstance, nil
// 		}
// 		fmt.Printf("Failed to connect to database. Retrying in %d seconds...\n", 1<<uint(i))
// 		time.Sleep(time.Second << uint(i))
// 	}

// 	return nil, nil, err
// }

func (cs *SQLConnection) Connect() (*gorm.DB, *sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", cs.username, cs.password, cs.host, cs.port, cs.database)
	Log.Info("Connection string: ", dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	Log.Info("Database connection successful!")
	dbInstance, _ := db.DB()

	return db, dbInstance, nil
}

func Close(db *sql.DB) (err error) {
	return db.Close()
}

func InitDBConnection(host string, port int, username string, password string) *SQLConnection {
	// Initialize MySQL connector
	connector := MySQLConnect(host, port, username, password, "open_weather")

	// Connect to MySQL
	ormdb, dbConn, err := connector.Connect()
	if err != nil {
		Log.Error("Error connecting to MySQL:", err)
	}

	return &SQLConnection{
		GormConn: ormdb,
		SqlCon:   dbConn,
	}
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
