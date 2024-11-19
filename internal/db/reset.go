package db

import (
	"STUOJ/internal/entity"
	"log"
)

func ResetDatabase() error {
	log.Println("Resetting database...")

	err := Db.AutoMigrate(&entity.Judgement{}, &entity.Language{}, &entity.Problem{}, &entity.History{}, &entity.ProblemTag{}, &entity.Solution{}, &entity.Submission{}, &entity.Tag{}, &entity.Testcase{}, &entity.User{})
	if err != nil {
		log.Println("Failed to migrate database!")
		return err
	}

	log.Println("Database reset success!")
	return nil
}
