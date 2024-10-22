package bootstrap

import (
	"STUOJ/conf"
)

func InitConfig() error {
	err := conf.InitConfig()
	if err != nil {
		return err
	}

	return nil
}
