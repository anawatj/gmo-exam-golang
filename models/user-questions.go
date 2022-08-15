package models

import (
	"database/sql"
	"strings"
)

type UserQuestion struct {
	ID                  string               `json:"id"`
	UserID              string               `json:"userId"`
	QuestionID          string               `json:"questionId"`
	Status              string               `json:"status"`
	UserName            string               `json:"userName"`
	QuestionDesc        string               `json:"questionDesc"`
	UserQuestionAnswers []UserQuestionAnswer `json:"userQuestionAnswers"`
}

type UserNameQuestion struct {
	UserName string `json:"userName"`
}

func (userQuestion *UserQuestion) Validate() []string {
	errors := []string{}
	if len(strings.TrimSpace(userQuestion.QuestionID)) == 0 {
		errors = append(errors, "Please input question")
	}
	if len(strings.TrimSpace(userQuestion.UserName)) == 0 {
		errors = append(errors, "Please input username")
	}
	return errors
}

func (userQuestion *UserQuestion) CreateUserQuestion(db *sql.DB) error {
	sql := `
		INSERT INTO public.user_questions
		(
			id,
			user_id,
			question_id,
			status
		)VALUES(
			$1,
			$2,
			$3,
			$4
		)
	`
	if _, err := db.Exec(sql, userQuestion.ID, userQuestion.UserID, userQuestion.QuestionID, userQuestion.Status); err != nil {
		return err
	}
	return nil

}
func (userQuestion *UserQuestion) UpdateUserQuestion(db *sql.DB) error {
	sql := `
		UPDATE public.user_questions SET
			user_id = $1,
			question_id = $2,
			status=$3
		WHERE id = $4
	`
	if _, err := db.Exec(sql, userQuestion.UserID, userQuestion.QuestionID, userQuestion.Status, userQuestion.ID); err != nil {
		return err
	}
	return nil
}

func (userQuestion *UserQuestion) GetUserQuestionByUser(db *sql.DB) ([]UserQuestion, error) {
	sql := `
		SELECT
			uq.id,
			uq.user_id,
			uq.question_id,
			uq.status,
			u.username, 
			q.question_desc 
		FROM public.user_questions uq 
		LEFT JOIN public.users u ON uq.user_id = u.id 
		LEFT JOIN public.questions q ON uq.question_id = q.id 
		WHERE u.username = $1
	`
	rows, err := db.Query(sql, userQuestion.UserName)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	userQuestions := []UserQuestion{}
	for rows.Next() {
		var userQuestion UserQuestion
		if err := rows.Scan(&userQuestion.ID, &userQuestion.UserID, &userQuestion.QuestionID, &userQuestion.Status, &userQuestion.UserName, &userQuestion.QuestionDesc); err != nil {
			return nil, err
		}
		userQuestions = append(userQuestions, userQuestion)
	}
	return userQuestions, nil
}
