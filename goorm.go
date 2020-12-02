package goorm

import (
	"database/sql"
	"goorm/log"
	"goorm/session"
)

// Engine is in charge of some preparations and finishing touchings,
// say connecting and testing DB before interaction and closing DB
// after interaction.
type Engine struct {
	db *sql.DB
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if err != nil {
		log.Error(err)
		return
	}

	// ping to DB to ensure that the DB connection is alive
	if err = db.Ping(); err != nil {
		log.Error(err)
		return
	}

	e = &Engine{db: db}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); err != nil {
		log.Error("Failed to close database")
	}
	log.Info("Close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.NewSession(e.db)
}
