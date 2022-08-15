package models

import "database/sql"

type Question struct {
	ID           string   `json:"id"`
	QuestionDesc string   `json:"questionDesc"`
	Answers      []Answer `json:"answers"`
}

func (question *Question) GetAllQuestion(db *sql.DB) ([]Question, error) {
	sql := `
		SELECT 
			id,
			question_desc
		FROM public.questions
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	questions := []Question{}
	for rows.Next() {
		var question Question
		if err := rows.Scan(&question.ID, &question.QuestionDesc); err != nil {
			return nil, err
		}
		question.Answers = []Answer{}
		questions = append(questions, question)
	}
	return questions, nil
}
