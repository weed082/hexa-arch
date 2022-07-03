package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Mysql struct {
	logger *log.Logger
	db     *sql.DB
}

func NewMysql(logger *log.Logger, driverName, dataSourceName string) *Mysql {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		logger.Fatalf("db connection failure: %v", err)
	}
	err = db.Ping()
	if err != nil {
		logger.Fatalf("db ping failure: %v", err)
	}
	return &Mysql{
		logger: logger,
		db:     db,
	}
}

// disconnect to mongoDB
func (sql *Mysql) Disconnect() {
	err := sql.db.Close()
	if err != nil {
		sql.logger.Fatalf("db close failure: %v", err)
	}
}
