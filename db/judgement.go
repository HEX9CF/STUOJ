package db

import (
	"STUOJ/model"
	"log"
)

// 插入评测结果
func InsertJudgement(s model.Judgement) (uint64, error) {
	sql := "INSERT INTO tbl_judgement (submission_id, testcase_id, time, memory, stdout, stderr, compile_output, message, status) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := Mysql.Prepare(sql)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	result, err := stmt.Exec(s.SubmissionId, s.TestcaseId, s.Time, s.Memory, s.Stdout, s.Stderr, s.CompileOutput, s.Message, s.Status)
	log.Println(sql, s.SubmissionId, s.TestcaseId, s.Time, s.Memory, s.Stdout, s.Stderr, s.CompileOutput, s.Message, s.Status)
	if err != nil {
		return 0, err
	}

	// 获取插入ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return uint64(id), nil
}

// 根据提交ID查询评测结果
func SelectJudgementsBySubmissionId(sid uint64) ([]model.Judgement, error) {
	sql := "SELECT id, submission_id, testcase_id, time, memory, stdout, stderr, compile_output, message, status FROM tbl_judgement WHERE submission_id = ?"
	rows, err := Mysql.Query(sql, sid)
	log.Println(sql, sid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	judgements := make([]model.Judgement, 0)
	for rows.Next() {
		var judgement model.Judgement

		err := rows.Scan(&judgement.Id, &judgement.SubmissionId, &judgement.TestcaseId, &judgement.Time, &judgement.Memory, &judgement.Stdout, &judgement.Stderr, &judgement.CompileOutput, &judgement.Message, &judgement.Status)
		if err != nil {
			return nil, err
		}

		//log.Println(judgement)
		judgements = append(judgements, judgement)
	}
	return judgements, nil
}

// 根据测试点ID查询评测结果
func SelectJudgementsByTestcaseId(tpid uint64) ([]model.Judgement, error) {
	sql := "SELECT id, submission_id, testcase_id, time, memory, stdout, stderr, compile_output, message, status FROM tbl_judgement WHERE testcase_id = ?"
	rows, err := Mysql.Query(sql, tpid)
	log.Println(sql, tpid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 遍历查询结果
	judgements := make([]model.Judgement, 0)
	for rows.Next() {
		var judgement model.Judgement

		err := rows.Scan(&judgement.Id, &judgement.SubmissionId, &judgement.TestcaseId, &judgement.Time, &judgement.Memory, &judgement.Stdout, &judgement.Stderr, &judgement.CompileOutput, &judgement.Message, &judgement.Status)
		if err != nil {
			return nil, err
		}

		//log.Println(judgement)
		judgements = append(judgements, judgement)
	}
	return judgements, nil
}

// 根据ID查询评测结果
func DeleteJudgementById(id uint64) error {
	sql := "DELETE FROM tbl_judgement WHERE id = ?"
	stmt, err := Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	log.Println(sql, id)
	if err != nil {
		return err
	}

	return nil
}

// 根据提交ID查询评测结果
func DeleteJudgementBySubmissionId(id uint64) error {
	sql := "DELETE FROM tbl_judgement WHERE submission_id = ?"
	stmt, err := Mysql.Prepare(sql)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)
	log.Println(sql, id)
	if err != nil {
		return err
	}

	return nil
}
