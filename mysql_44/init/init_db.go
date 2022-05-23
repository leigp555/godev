package init

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

var db *sql.DB

func init() {
	d, err := sql.Open("mysql", "root:123456@/testsql")
	if err != nil {
		log.Panic(err)
	}
	d.SetConnMaxLifetime(time.Minute * 3)
	d.SetMaxOpenConns(10)
	d.SetMaxIdleConns(10)
	db = d
}
func initDb() {
	fmt.Println(db)
}
