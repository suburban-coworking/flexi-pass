package enviro

import (
	"database/sql"
	"os"
	"path"

	"github.com/gobuffalo/packr"
	migrate "github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
)

type env struct {
	AppDir string
	DB     *sql.DB
	Log    *log.Logger
}

// Env gloabl cache of application environment.
var Env env

func (e *env) Init(dir, level string) {
	// Initialise logger
	e.setupLog(level)

	// Create application directory
	e.setupDir(dir)

	// Initialise database
	e.initDB(dir)
}

func (e *env) setupLog(level string) {
	e.Log = log.New()

	lvl, err := log.ParseLevel(level)
	if err != nil {
		e.Log.WithFields(log.Fields{
			"level": level,
		}).Fatal("Invalid logging level!")
	}

	e.Log.SetLevel(lvl)
	e.Log.WithFields(log.Fields{
		"level": lvl,
	}).Debug("Logger configured.")
}

func (e *env) setupDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		e.Log.WithFields(log.Fields{
			"directory": dir,
		}).Fatal("Failed to create application directory: ", err)
	}

	e.AppDir = dir
	e.Log.WithFields(log.Fields{
		"directory": Env.AppDir,
	}).Debug("Setup done.")
}

func (e *env) initDB(dir string) {
	var err error

	dir = path.Join(dir, "su.db")

	e.Log.WithFields(log.Fields{
		"db": dir,
	}).Debug("Opening database connection pool.")

	e.DB, err = sql.Open("sqlite3", dir)

	if err != nil {
		e.Log.Fatal("Failed to open database: ", err)
	}

	defer e.DB.Close()

	e.migrateDB()
}

func (e *env) migrateDB() {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("../migrations"),
	}

	n, err := migrate.Exec(e.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		e.Log.Fatal("Failed to migrate database: ", err)
	}

	e.Log.WithFields(log.Fields{
		"count": n,
	}).Debug("DB migrations applied.")
}
