package migration

import (
	"fmt"
	"time"

	database "assesment/config"
	entities "assesment/entities"

	"github.com/google/uuid"
)

func Seeder()error{
	db := database.SetUpDatabaseConnection()
	db.AutoMigrate(&entities.Answer{}, &entities.Assessment{}, &entities.Choice{}, &entities.Question{}, &entities.Submission{})

	for i := 1; i <= 5; i++ {
		assessmentID := uuid.MustParse(fmt.Sprintf("00000000-0000-0000-0000-00000000000%d", i))
		classID := uuid.MustParse(fmt.Sprintf("%d%d%d%d%d%d%d%d-%d%d%d%d-%d%d%d%d-%d%d%d%d-%d%d%d%d%d%d%d%d%d%d%d%d",i, i, i, i, i, i, i, i,i, i, i, i,i, i, i, i,i, i, i, i,i, i, i, i, i, i, i, i, i, i, i, i))
		now := time.Now()
		assessment := entities.Assessment{
			ID:        assessmentID,
			Name:      fmt.Sprintf("Ujian %d", i),
			StartTime: now,
			EndTime:   now.Add(time.Hour),
			Duration:  3600,
			ClassID:   classID,
			CreatedAt: now,
			UpdatedAt: now,
		}
		db.Create(&assessment)

		// Seed Questions
		questionID := uuid.MustParse(fmt.Sprintf("10000000-0000-0000-0000-00000000000%d", i))
		question := entities.Question{
			ID:           questionID,
			QuestionText: fmt.Sprintf("Apa itu soal nomor %d?", i),
			EvaluationID: assessmentID,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		db.Create(&question)

		// Seed Choices
		for j := 1; j <= 4; j++ {
			choiceID := uuid.MustParse(fmt.Sprintf("20000000-0000-0000-0000-0000000000%d%d", i, j))
			choice := entities.Choice{
				ID:         choiceID,
				ChoiceText: fmt.Sprintf("Pilihan %d untuk soal %d", j, i),
				QuestionID: questionID,
				IsCorrect:  j == 1,
				CreatedAt:  now,
				UpdatedAt:  now,
			}
			db.Create(&choice)
		}

		// Seed Submissions
		submissionID := uuid.MustParse(fmt.Sprintf("30000000-0000-0000-0000-00000000000%d", i))
		userID := uuid.MustParse(fmt.Sprintf("%d%d%d%d%d%d%d%d-%d%d%d%d-%d%d%d%d-%d%d%d%d-%d%d%d%d%d%d%d%d%d%d%d%d",i, i, i, i, i, i, i, i+1,i+1, i+1, i+1, i+2,i+2, i+2, i+2, i+3,i+3, i+3, i+3, i+4,i+4, i+4, i+4, i+4, i+4, i+4, i+4, i+4, i+4, i+4, i+4, i+4))
		submission := entities.Submission{
			ID:           submissionID,
			UserID:       userID,
			AssessmentID: assessmentID,
			EndedTime:    now.Add(30 * time.Minute),
			SubmittedAt:  now.Add(30 * time.Minute),
			Score:        100,
			Status:       entities.StatusSubmitted,
		}
		db.Create(&submission)
		
		// Seed Answer
		answerID := uuid.MustParse(fmt.Sprintf("50000000-0000-0000-0000-00000000000%d", i))
		choiceID := uuid.MustParse(fmt.Sprintf("20000000-0000-0000-0000-0000000000%d1", i)) // asumsi pilihan pertama yang benar
		answer := entities.Answer{
			ID:           answerID,
			QuestionID:   questionID,
			ChoiceID:     choiceID,
			SubmissionID: submissionID,
			CreatedAt:    now,
			UpdatedAt:    now,
		}
		db.Create(&answer)
	}

	fmt.Println("Seeder selesai.")
	return nil
}