package models

import (
	"database/sql"
	"strings"
)

type UserQuestionAnswer struct {
	ID             string `json:"id"`
	UserQuestionID string `json:"userQuestionId"`
	AnswerID       string `json:"answerId"`
	AnswerDesc     string `json:"answerDesc"`
	AnswerScore    int    `json:"answerScore"`
}

func (userQuestionAnswer *UserQuestionAnswer) Validate() []string {
	errors := []string{}
	if len(strings.TrimSpace(userQuestionAnswer.AnswerID)) == 0 {
		errors = append(errors, "Please input answers")
	}
	return errors
}
func (userQuestionAnswer *UserQuestionAnswer) CreateUserQuestionAnswer(db *sql.DB) error {
	sql := `
		INSERT INTO public.user_question_answers
		(
			id,
			user_question_id,
			answer_id
		)VALUES(
			$1,
			$2,
			$3
		)
	`
	if _, err := db.Exec(sql, userQuestionAnswer.ID, userQuestionAnswer.UserQuestionID, userQuestionAnswer.AnswerID); err != nil {
		return err
	}
	return nil
}
func (userQuestionAnswwer *UserQuestionAnswer) DeleteUserQuestionAnswersByUserQuestion(db *sql.DB) error {
	sql := `
		DELETE 
		FROM public.user_question_answers
		WHERE user_question_id = $1
	`
	_, err := db.Exec(sql, userQuestionAnswwer.UserQuestionID)
	if err != nil {
		return err
	}
	return nil
}
func (userQuestionAnswer *UserQuestionAnswer) GetUserQuestionAnswerByUserQuestion(db *sql.DB) ([]UserQuestionAnswer, error) {
	sql := `
		SELECT 
			uqa.id ,
			uqa.user_question_id,
			uqa.answer_id ,
			a.answer_desc ,
			a.answer_score 
		FROM public.user_question_answers uqa 
		LEFT JOIN public.answers a ON uqa.answer_id = a.id 
		WHERE uqa.user_question_id = $1
	`
	rows, err := db.Query(sql, userQuestionAnswer.UserQuestionID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	userQuestionAnswers := []UserQuestionAnswer{}
	for rows.Next() {
		var userQuestionAnswer UserQuestionAnswer
		if err := rows.Scan(&userQuestionAnswer.ID, &userQuestionAnswer.UserQuestionID, &userQuestionAnswer.AnswerID, &userQuestionAnswer.AnswerDesc, &userQuestionAnswer.AnswerScore); err != nil {
			return nil, err
		}
		userQuestionAnswers = append(userQuestionAnswers, userQuestionAnswer)
	}
	return userQuestionAnswers, nil
}
