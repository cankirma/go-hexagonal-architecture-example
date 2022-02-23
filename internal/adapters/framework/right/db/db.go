package db

import "database/sql"

//import (
//	"database/sql"
//	"log"
//	"time"
// sq	"github.com/Masterminds/squirrel"
//	_ "github.com//go-sql-driver/mysql"
//)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	//connect
	db, err := sql.Open(driverName, dataSourceName)
}
