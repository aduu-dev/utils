package helper

import (
	"database/sql"

	"github.com/pkg/errors"
)

// CleanupWithError rolls the given transaction back and if that returns errors it combines the given error and the new.
func CleanupWithError(tx *sql.Tx, err error) error {
	if err2 := tx.Rollback(); err2 != nil {
		return errors.Wrapf(err, "failed to cleanup after: %v", err2)
	}
	return err
}
