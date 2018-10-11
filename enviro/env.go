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

func (e *env) Init(dir, level string) error {
	var err error

	// Initialise logger
	if err = e.setupLog(level); err != nil {
		return err
	}

	// Create application directory
	if err = e.setupDir(dir); err != nil {
		return err
	}

	// Initialise database
	if err = e.initDB(dir); err != nil {
		return err
	}
	return nil
}

func (e *env) setupLog(level string) error {
	e.Log = log.New()
	lvl, err := log.ParseLevel(level)

	if err != nil {
		return err
	}
	e.Log.SetLevel(lvl)

	e.Log.WithFields(log.Fields{
		"level": lvl,
	}).Debug("Logger configured.")
	return nil
}

func (e *env) setupDir(dir string) error {
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		return err
	}
	e.AppDir = dir

	e.Log.WithFields(log.Fields{
		"directory": Env.AppDir,
	}).Debug("Setup done.")
	return nil
}

func (e *env) initDB(dir string) error {
	var err error

	dir = path.Join(dir, "su.db")
	e.DB, err = sql.Open("sqlite3", dir)

	if err != nil {
		return err
	}
	defer e.DB.Close()

	e.Log.WithFields(log.Fields{
		"db": dir,
	}).Debug("DB pool opened, migrating database.")
	return e.migrateDB()
}

func (e *env) migrateDB() error {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("../migrations"),
	}

	n, err := migrate.Exec(e.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		return err
	}

	e.Log.WithFields(log.Fields{
		"count": n,
	}).Debug("DB migrations applied.")
	return nil
}
