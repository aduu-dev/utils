package dbhelper

import (
	"github.com/jmoiron/sqlx"
)

func CreateSchema(db *sqlx.DB, schema []string) error {
	// Initializing sqlite db.
	for _, subSchema := range schema {
		_, err := db.Exec(subSchema)
		if err != nil {
			db.Close()
			return err
		}
	}
	return nil
}
