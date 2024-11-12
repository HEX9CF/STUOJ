package judge

import (
	"STUOJ/external/judge0"
	"STUOJ/internal/entity"
	"STUOJ/internal/model"
	"STUOJ/internal/service/problem"
	"STUOJ/internal/service/record"
	"STUOJ/internal/service/testcase"
	"errors"
	"log"
	"math"
	"strconv"
)

func Submit(s entity.Submission) (uint64, error) {
	var err error

	// 获取代码长度
	s.Length = uint64(len(s.SourceCode))

	// 获取题目信息
	p, err := problem.SelectProblemById(s.ProblemId)
	if err != nil {
		log.Println(err)
		return 0, errors.New("获取题目信息失败")
	}

	// 获取评测点
	testcases, err := testcase.SelectByProblemId(s.ProblemId)
	if err != nil {
		log.Println(err)
		return 0, errors.New("获取评测点数据失败")
	}

	// 插入提交
	s.Id, err = record.InsertSubmission(s)
	if err != nil {
		log.Println(err)
		return 0, errors.New("插入提交信息失败")
	}

	// 异步提交
	go asyncSubmit(s, p, testcases)

	return s.Id, nil
}

// 异步提交
func asyncSubmit(s entity.Submission, p entity.Problem, ts []entity.Testcase) {
	s.Status = entity.SubmitStatusAC
	chJudgement := make(chan entity.Judgement)

	// 提交评测点
	for _, t := range ts {
		// 异步评测
		go asyncJudge(s, p, t, chJudgement)
	}

	for _, _ = range ts {
		// 接收评测点结果
		j := <-chJudgement
		//log.Println(j)

		// 更新评测点结果
		err := record.UpdateJudgementById(j)
		if err != nil {
			log.Println(err)
			return
		}

		// 更新提交数据
		s.Time = math.Max(s.Time, j.Time)
		s.Memory = max(s.Memory, j.Memory)
		// 如果评测点结果不是AC，更新提交状态
		if j.Status != entity.SubmitStatusAC {
			if s.Status != entity.SubmitStatusWA {
				s.Status = max(s.Status, j.Status)
			}
		}
	}

	// 更新提交信息
	err := record.UpdateSubmissionById(s)
	if err != nil {
		log.Println(err)
		return
	}
}

// 异步评测
func asyncJudge(s entity.Submission, p entity.Problem, t entity.Testcase, ch chan entity.Judgement) {
	var err error

	// 初始化评测点结果对象
	j := entity.Judgement{
		SubmissionId: s.Id,
		TestcaseId:   t.Id,
		Status:       entity.SubmitStatusPend,
	}

	// 插入评测点结果
	j.Id, err = record.InsertJudgement(j)
	if err != nil {
		log.Println(err)
		ch <- j
		return
	}

	// 初始化评测点评测对象
	judgeSubmission := model.JudgeSubmission{
		SourceCode:     s.SourceCode,
		LanguageId:     s.LanguageId,
		Stdin:          t.TestInput,
		ExpectedOutput: t.TestOutput,
		CPUTimeLimit:   p.TimeLimit,
		MemoryLimit:    p.MemoryLimit,
	}
	//log.Println(judgeSubmission)

	// 发送评测点评测请求（等待评测结果）
	result, err := judge0.Submit(judgeSubmission)
	if err != nil {
		log.Println(err)
		j.Status = entity.SubmitStatusIE
		ch <- j
		return
	}
	//log.Println(result)

	// 解析时间
	time := float64(0)
	if result.Time != "" {
		time, err = strconv.ParseFloat(result.Time, 64)
		if err != nil {
			log.Println(err)
			j.Status = entity.SubmitStatusIE
			ch <- j
			return
		}
	}

	// 更新评测点结果
	j.Time = time
	j.Memory = uint64(result.Memory)
	j.Stdout = result.Stdout
	j.Stderr = result.Stderr
	j.CompileOutput = result.CompileOutput
	j.Message = result.Message
	j.Status = entity.SubmitStatus(result.Status.Id)
	//log.Println(j)

	// 发送评测点结果到通道
	ch <- j
}

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

	time, err := strconv.ParseFloat(res.Time, 64)
	if err != nil {
		log.Println(err)
		return entity.Judgement{}, errors.New("时间解析失败")
	}

	j := entity.Judgement{
		Stdout: res.Stdout,
		Time:   time,
		Memory: uint64(res.Memory),
	}

	return j, nil
}
