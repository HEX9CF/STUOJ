package language

import "STUOJ/internal/dao"

// 删除所有语言
func DeleteAll() error {
	err := dao.DeleteAllLanguages()
	if err != nil {
		return err
	}

	return nil
}
