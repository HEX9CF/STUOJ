package bootstrap

import (
	"STUOJ/handlers"
)

func InitHandlers() error {
	err := handlers.InitHandlers()
	if err != nil {
		return err
	}

	return nil
}
