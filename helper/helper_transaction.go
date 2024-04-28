package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	if r := recover(); r != nil {
		rollbackError := tx.Rollback()
		HelperPanic(rollbackError)
	} else {
		commitError := tx.Commit()
		HelperPanic(commitError)
	}
}