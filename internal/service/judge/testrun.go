package judge

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"errors"
	"log"
	"strconv"
)

func TestRun(s entity.Submission, stdin string) (entity.Judgement, error) {
	js := model.JudgeSubmission{
		SourceCode: s.SourceCode,
		LanguageId: s.LanguageId,
		Stdin:      stdin,
	}

	res, err := judge0.Submit(js)
	if err != nil {
		log.Println(err)
		return entity.Judgement{}, errors.New("提交失败")
	}

	// 解析时间
	time := float64(0)
	if res.Time != "" {
		time, err = strconv.ParseFloat(res.Time, 64)
		if err != nil {
			log.Println(err)
			return entity.Judgement{}, errors.New("时间解析失败")
		}
	}

	j := entity.Judgement{
		Stdout: res.Stdout,
		Time:   time,
		Memory: uint64(res.Memory),
	}

	return j, nil
}
