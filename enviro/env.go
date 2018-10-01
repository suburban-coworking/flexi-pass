package enviro

import (
	"database/sql"
	// blank import 'sqlite3' as we want to rely solely on the 'sql' package's interface
	_ "github.com/mattn/go-sqlite3"

	log "github.com/sirupsen/logrus"
)

type env struct {
	DB  *sql.DB
	Log *log.Logger
}

// Env - applications current environment.  This is quite limited at the
// moment, exposing the database pool and logger to the rest of the
// application.
var Env env

// Init the application's environment.
// This involves setting up the logger and opening our connection pool to the database.
func Init(dsn, lvl string) {
	// Initialise the logger
	initLog(lvl)

	// Initialise database connection pool
	initDb(dsn)
}

func initLog(lvl string) {
	Env.Log = log.New()
	// Then convert the string level to a useful logging level
	level, err := log.ParseLevel(lvl)
	if err != nil {
		log.Fatal("Failed to initialise logger: ", err)
	}

	Env.Log.SetLevel(level)
}

func initDb(dsn string) {
	var err error
	Env.Log.WithFields(log.Fields{
		"db": dsn,
	}).Debug("Opening database connection pool.")

	Env.DB, err = sql.Open("sqlite3", dsn)

	if err != nil {
		Env.Log.Fatal("Failed to open database: ", err)
	}

	defer Env.DB.Close()
}
