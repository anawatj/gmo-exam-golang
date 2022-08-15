package models

import (
	"database/sql"
)

type Answer struct {
	ID          string `json:"id"`
	QuestionID  string `json:"questionId"`
	AnswerDesc  string `json:"answerDesc"`
	AnswerScore int    `json:"answerScore"`
}

func (answer *Answer) GetAnswerByQuestion(db *sql.DB) ([]Answer, error) {
	sql := `
		SELECT id,question_id,answer_desc,answer_score
		FROM public.answers
		WHERE question_id = $1
	`
	rows, err := db.Query(sql, answer.QuestionID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	answers := []Answer{}
	for rows.Next() {
		var answer Answer
		if err := rows.Scan(&answer.ID, &answer.QuestionID, &answer.AnswerDesc, &answer.AnswerScore); err != nil {
			return nil, err
		}
		answers = append(answers, answer)
	}
	return answers, nil
}
