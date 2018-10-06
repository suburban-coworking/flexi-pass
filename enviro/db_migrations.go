package enviro

import (
	"github.com/gobuffalo/packr"
	"github.com/rubenv/sql-migrate"
	log "github.com/sirupsen/logrus"
)

func migrateDB() {
	log.Debug("Upgrading database ...")
	migrations := &migrate.PackrMigrationSource{
		Box: packr.NewBox("./migrations"),
	}

	n, err := migrate.Exec(Env.DB, "sqlite3", migrations, migrate.Up)
	if err != nil {
		log.Fatal("Failed to execute migrations.\n", err)
	}

	log.WithFields(log.Fields{
		"count": n,
	}).Debug("Migrations applied.")
}
