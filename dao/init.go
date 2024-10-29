package dao

import (
	"database/sql"
	"os"
)

func (d *DAO) InitSQL(initSQLFile string) error {
	var db *sql.DB = (*sql.DB)(d)

	var (
		bytes []byte
		err   error
	)

	if bytes, err = os.ReadFile(initSQLFile); err != nil {
		return err
	}

	if _, err = db.Exec(string(bytes)); err != nil {
		return err
	}
	return nil
}
