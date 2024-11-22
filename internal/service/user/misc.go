package user

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"log"
	"time"
)

// 统计用户
func GetStatistics(p model.Period) (model.UserStatistics, error) {
	var err error
	var cbds []model.CountByDate
	var cbrs []model.CountByRole
	var stats model.UserStatistics

	// 检查时间范围
	if p.StartTime.After(p.EndTime) {
		return model.UserStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计用户数量
	stats.UserCount, err = dao.CountUsers()
	if err != nil {
		log.Println(err)
		return model.UserStatistics{}, errors.New("统计用户数量失败")
	}

	// 统计用户角色
	cbrs, err = dao.CountUsersGroupByRole()
	if err != nil {
		log.Println(err)
		return model.UserStatistics{}, errors.New("统计用户角色失败")
	}
	stats.UserCountByRole.FromCountByRole(cbrs)
	// 统计用户注册数量
	cbds, err = dao.CountUsersBetweenCreateTime(p.StartTime, p.EndTime)
	if err != nil {
		log.Println(err)
		return model.UserStatistics{}, errors.New("统计用户注册数量失败")
	}
	stats.RegisterCountByDate.FromCountByDate(cbds)
	fillZero(&stats, p.StartTime, p.EndTime)

	return stats, nil
}

func fillZero(s *model.UserStatistics, startDate time.Time, endDate time.Time) {
	utils.MapCountFillZero(&s.UserCountByRole, startDate, endDate)
	utils.MapCountFillZero(&s.RegisterCountByDate, startDate, endDate)
}
