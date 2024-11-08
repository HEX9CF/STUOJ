package user

import "STUOJ/internal/db/dao"

// 根据ID删除用户
func DeleteById(id uint64) error {
	err := dao.DeleteUserById(id)
	if err != nil {
		return err
	}

	return nil
}
