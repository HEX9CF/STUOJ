package problem

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/utils"
	"errors"
	"log"
	"time"
)

// 提交记录统计
func GetStatistics(p model.Period) (model.ProblemStatistics, error) {
	var err error
	var cbds []model.CountByDate
	var stats model.ProblemStatistics

	// 检查时间范围
	if p.StartTime.After(p.EndTime) {
		return model.ProblemStatistics{}, errors.New("开始时间不能晚于结束时间")
	}

	// 统计题目数量
	stats.ProblemCount, err = dao.CountProblems()
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计题目数量失败")
	}

	// 统计评测点数量
	stats.TestcaseCount, err = dao.CountTestcases()
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计评测点数量失败")
	}

	// 统计题解数量
	stats.SolutionCount, err = dao.CountBlogs()
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计题解数量失败")
	}

	// 统计添加题目数量
	cbds, err = dao.CountHistoriesBetweenCreateTimeByOperation(p.StartTime, p.EndTime, entity.OperationInsert)
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计添加题目数量失败")
	}
	stats.InsertCountByDate.FromCountByDate(cbds)

	// 统计修改题目数量
	cbds, err = dao.CountHistoriesBetweenCreateTimeByOperation(p.StartTime, p.EndTime, entity.OperationUpdate)
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计修改题目数量失败")
	}
	stats.UpdateCountByDate.FromCountByDate(cbds)

	// 统计删除题目数量
	cbds, err = dao.CountHistoriesBetweenCreateTimeByOperation(p.StartTime, p.EndTime, entity.OperationDelete)
	if err != nil {
		log.Println(err)
		return model.ProblemStatistics{}, errors.New("统计删除题目数量失败")
	}
	stats.DeleteCountByDate.FromCountByDate(cbds)
	fillZero(&stats, p.StartTime, p.EndTime)

	return stats, nil
}

func fillZero(p *model.ProblemStatistics, startDate time.Time, endDate time.Time) {
	utils.MapCountFillZero(&p.InsertCountByDate, startDate, endDate)
	utils.MapCountFillZero(&p.UpdateCountByDate, startDate, endDate)
	utils.MapCountFillZero(&p.DeleteCountByDate, startDate, endDate)
}
