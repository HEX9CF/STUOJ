package bootstrap

import (
	"STUOJ/db"
)

func InitDatabase() error {
	err := db.InitDatabase()
	if err != nil {
		return err
	}

	return nil
}
