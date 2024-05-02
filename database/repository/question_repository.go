package repository

import (
	"gobot/database"
	"gobot/models"
)

func GetQuestions() ([]models.Question, error) {
	rows, err := database.DB.Query("SELECT * FROM question")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []models.Question
	for rows.Next() {
		var question models.Question
		err := rows.Scan(&question.ID, &question.Text)
		if err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}

	return questions, nil
}
