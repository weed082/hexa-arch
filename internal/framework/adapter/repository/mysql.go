package repository

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	db *sql.DB
}

func NewMysql(driverName, dataSourceName string) *Mysql {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failur: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}
	return &Mysql{db: db}
}

// disconnect to mongoDB
func (mysql Mysql) Disconnect() {
	err := mysql.db.Close()
	if err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}
