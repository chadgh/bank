package scripts

import "database/sql"

func TruncateAll(db *sql.DB) {
	db.Exec("TRUNCATE account_transactions;")
}
