package models

import (
	"database/sql"
	"strings"
)

type User struct {
	ID          string `json:"id"`
	UserName    string `json:"userName"`
	UserGroupID string `json:"userGroupId"`
}

type Summary struct {
	UserName string `json:"userName"`
	Score    int    `json:"score"`
	Rank     int    `json:"rank"`
}

func (user *User) ValidateUser() []string {
	var errors []string
	if strings.TrimSpace(user.UserGroupID) == "" {
		errors = append(errors, "User group is required")
	}
	if strings.TrimSpace(user.UserName) == "" {
		errors = append(errors, "Username is required")
	}
	return errors
}
func (user *User) GetUserByUserName(db *sql.DB) error {
	sql := `
		SELECT 
			id,
			username,
			user_group_id
		FROM
			public.users
		WHERE username = $1
		limit 1 
	`
	return db.QueryRow(sql, user.UserName).Scan(&user.ID, &user.UserName, &user.UserGroupID)
}
func (user *User) CreateUser(db *sql.DB) error {
	sql := `
		INSERT INTO public.users
		(
			id,
			username,
			user_group_id
		)VALUES(
			$1,
			$2,
			$3
		)
	`
	_, err := db.Exec(sql, user.ID, user.UserName, user.UserGroupID)
	if err != nil {
		return err
	}
	return nil
}
func (user *User) Summary(db *sql.DB) (Summary, error) {
	sql := `
		SELECT 
			u.username ,
			SUM(a.answer_score) As Score
		FROM
			public.users u 
			JOIN public.user_questions uq ON u.id = uq.user_id 
			JOIN public.user_question_answers uqa ON uq.id= uqa.user_question_id
			JOIN public.answers a ON uqa.answer_id = a.id 
		WHERE u.username = $1
		GROUP BY u.username 
	`
	rows, err := db.Query(sql, user.UserName)
	if err != nil {
		return Summary{}, err
	}

	defer rows.Close()
	summaries := []Summary{}
	for rows.Next() {
		var summary Summary
		if err := rows.Scan(&summary.UserName, &summary.Score); err != nil {
			return Summary{}, err
		}
		summaries = append(summaries, summary)
	}
	return summaries[0], nil

}
