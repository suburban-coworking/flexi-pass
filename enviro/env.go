package enviro

import (
	"database/sql"
	"os"

	log "github.com/sirupsen/logrus"
)

type env struct {
	AppDir string
	DB     *sql.DB
}

// Env gloabl cache of application environment.
var Env env

func (e *env) Init(dir string) {
	// Create application directory
	setupDir(dir)
}

func setupDir(dir string) {
	err := os.MkdirAll(dir, os.ModePerm)

	if err != nil {
		log.WithFields(log.Fields{
			"directory": dir,
		}).Fatal("Failed to create application directory.", err)
	}

	Env.AppDir = dir
	log.WithFields(log.Fields{
		"directory": Env.AppDir,
	}).Info("Application directory set.")
}
