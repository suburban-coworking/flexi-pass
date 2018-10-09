package enviro

import (
	"database/sql"
	"os"

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
	setupDir(dir)
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

func setupDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.WithFields(log.Fields{
			"directory": dir,
		}).Fatal("Failed to create application directory!", err)
	}

	Env.AppDir = dir
	log.WithFields(log.Fields{
		"directory": Env.AppDir,
	}).Debug("Setup done.")
}
