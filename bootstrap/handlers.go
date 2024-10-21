package bootstrap

import (
	"STUOJ/handlers"
)

func InitHandlers() error {
	err := handlers.Init()
	if err != nil {
		return err
	}

	return nil
}
