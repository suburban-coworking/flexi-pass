package enviro

import (
	"database/sql"
	"os"
	"path"

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
func Init(dir, lvl string) {
	// Initialise the logger
	initLog(lvl)

	// Setup application directories.
	setupAppDir(dir)

	// Initialise database connection pool
	initDb(dir)
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

func setupAppDir(dir string) {
	Env.Log.Debug("Setting up application ...")

	err := os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatal("Failed to setup application directory.\n", err)
	}

	Env.Log.WithFields(log.Fields{
		"dir": dir,
	}).Debug("Application directory created.")
}

func initDb(dir string) {
	var err error

	dir = path.Join(dir, "su.db")

	Env.Log.WithFields(log.Fields{
		"db": dir,
	}).Debug("Opening database connection pool.")

	Env.DB, err = sql.Open("sqlite3", dir)

	if err != nil {
		Env.Log.Fatal("Failed to open database: ", err)
	}

	defer Env.DB.Close()

	migrateDB()
}
