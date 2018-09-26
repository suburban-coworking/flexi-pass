package enviro

import (
	"database/sql"

	log "github.com/sirupsen/logrus"
)

type env struct {
	db  *sql.DB
	log *log.Logger
}

// Init intialises the applications environment.
// This involves setting up the logger and opening our connection pool to the database.
func (e *env) Init(url, levelStr string) {
	// First, initialise the logger
	e.log = log.New()

	// Then convert the string level to a useful logging level
	level, err := log.ParseLevel(levelStr)
	if err != nil {
		log.Fatal("Failed to initialise logger: ", err)
	}

	e.log.SetLevel(level)

	e.db, err = sql.Open("sqlite3", url)
	if err != nil {
		log.Fatal("Failed to open db connection pool: ", err)
	}

	defer e.db.Close()
}
