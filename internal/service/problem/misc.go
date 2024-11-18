package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
	"time"
)

// 提交记录统计
func GetStatistics(startTime time.Time, endTime time.Time) (model.ProblemStatistics, error) {
	var stats model.ProblemStatistics

	// 检查时间范围
	if startTime.After(endTime) {
		return model.ProblemStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计添加题目数量
	cbds, err := dao.CountProblemHistoriesBetweenCreateTimeByOperation(startTime, endTime, entity.OperationInsert)
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计添加题目数量失败")
	}
	stats.InsertCountByDate.FromStruct(cbds)

	// 统计修改题目数量
	cbds, err = dao.CountProblemHistoriesBetweenCreateTimeByOperation(startTime, endTime, entity.OperationUpdate)
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计修改题目数量失败")
	}
	stats.UpdateCountByDate.FromStruct(cbds)

	// 统计删除题目数量
	cbds, err = dao.CountProblemHistoriesBetweenCreateTimeByOperation(startTime, endTime, entity.OperationDelete)
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计删除题目数量失败")
	}
	stats.DeleteCountByDate.FromStruct(cbds)

	return stats, nil
}
