package main

import "database/sql"

type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		Error(err)
		return
	}
	if err = db.Ping(); err != nil {
		Error(err)
		return
	}
	e = &Engine{db: db}
	Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		Error("Failed to close database")
	}
	Info("Close database success")
}

func (e *Engine) NewSession() *Session {
	return New(e.db)
}
