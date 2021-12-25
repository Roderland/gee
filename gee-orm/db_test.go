package gee_orm

import (
	"database/sql"
	"testing"
)

import (
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func TestDB(t *testing.T) {
	db, _ := sql.Open("sqlite3", "gee.db")
	defer func() { _ = db.Close() }()
	_, _ = db.Exec("drop table if exists User;")
	_, _ = db.Exec("create table User(Name text);")
	result, err := db.Exec("insert into User (`Name`) values (?), (?);", "Tom", "Jack")
	if err == nil {
		affected, _ := result.RowsAffected()
		log.Println(affected)
	}
	row := db.QueryRow("select Name from User limit 1;")
	var name string
	if err = row.Scan(&name); err == nil {
		log.Println(name)
	}
}
