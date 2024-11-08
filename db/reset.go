package db

import (
	"STUOJ/model"
	"log"
)

func ResetDatabase() error {
	log.Println("Resetting database...")

	err := Db.AutoMigrate(&model.Judgement{}, &model.Language{}, &model.Problem{}, &model.ProblemHistory{}, &model.ProblemTag{}, &model.Solution{}, &model.Submission{}, &model.Tag{}, &model.Testcase{}, &model.User{})
	if err != nil {
		log.Println("Failed to migrate database!")
		return err
	}

	log.Println("Database reset success!")
	return nil
}
