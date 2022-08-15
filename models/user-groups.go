package models

import "database/sql"

type UserGroup struct {
	ID            string `json:"id"`
	UserGroupName string `json:"userGroupName"`
}

func (userGroup *UserGroup) GetAllUserGroup(db *sql.DB) ([]UserGroup, error) {
	sql := `
		SELECT id,user_group_name
		FROM public.user_groups
	`
	rows, err := db.Query(sql)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	userGroups := []UserGroup{}
	for rows.Next() {
		var userGroup UserGroup
		if err := rows.Scan(&userGroup.ID, &userGroup.UserGroupName); err != nil {
			return nil, err
		}
		userGroups = append(userGroups, userGroup)
	}
	return userGroups, nil

}
