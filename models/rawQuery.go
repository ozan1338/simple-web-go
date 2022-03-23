package models

func QueryUpdateUser(user User) string {
	sqlQuery := "update users set "
	if user.FirstName != "" {
		sqlQuery += "first_name = '" + user.FirstName + "' "
	}

	if user.LastName != "" {
		sqlQuery += "last_name = '" + user.LastName + "' "
	}

	if user.Email != "" {
		sqlQuery += "email = '" + user.Email + "' "
	}

	where := "where id = ? returning *"

	sqlQuery += where

	return sqlQuery
}