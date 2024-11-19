package tag

import (
	"STUOJ/internal/dao"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
)

func GetStatistics() (model.TagStatistics, error) {
	var err error
	var stats model.TagStatistics
	var cbts []model.CountByTag

	// 统计标签数量
	stats.TagCount, err = dao.CountTags()
	if err != nil {
		log.Println(err)
		return model.TagStatistics{}, errors.New("统计标签数量失败")
	}

	// 统计题目数量
	cbts, err = dao.CountProblemsGroupByTag()
	if err != nil {
		log.Println(err)
		return model.TagStatistics{}, errors.New("统计题目数量失败")
	}
	stats.ProblemCountByTag = make(model.MapCount)
	for _, v := range cbts {
		var tag entity.Tag
		tag, err = dao.SelectTagById(v.TagId)
		if err != nil {
			log.Println(err)
			tag = entity.Tag{
				Id:   v.TagId,
				Name: "未知标签",
			}
		}
		stats.ProblemCountByTag[tag.Name] = v.Count
	}

	return stats, nil
}
