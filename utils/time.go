package utils

import (
	"STUOJ/internal/model"
	"github.com/gin-gonic/gin"
	"time"
)

const (
	DATETIME_LAYOUT = "2006-01-02 15:04:05"
	DATE_LAYOUT     = "2006-01-02"
)

func GenerateDateList(startDate time.Time, endDate time.Time) []string {
	var dateList []string
	for startDate.Before(endDate.AddDate(0, 0, 1)) {
		dateList = append(dateList, startDate.Format(DATE_LAYOUT))
		startDate = startDate.AddDate(0, 0, 1)
	}
	return dateList
}

func MapCountFillZero(m *model.MapCount, startDate time.Time, endDate time.Time) {
	dateList := GenerateDateList(startDate, endDate)
	// 填充没有结果的日期
	for _, date := range dateList {
		if _, ok := (*m)[date]; !ok {
			(*m)[date] = 0
		}
	}
}

// 从请求中获取时间范围
func GetPeriod(c *gin.Context) (model.Period, error) {
	var err error
	var p model.Period

	// 读取参数
	startTimeStr := c.Query("start-time")
	endTimeStr := c.Query("end-time")

	// 解析时间范围
	err = p.FromString(startTimeStr, endTimeStr, DATETIME_LAYOUT)
	if err != nil {
		return model.Period{}, err
	}

	// 检查时间范围
	err = p.Check()
	if err != nil {
		return model.Period{}, err
	}

	return p, nil
}
