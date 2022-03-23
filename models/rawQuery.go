package models

import "strconv"

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

	if user.RoleId != 0 {
		sqlQuery += "role_id = '"+ strconv.Itoa(int(user.RoleId)) + "' "
	}

	where := "where id = ? returning *"

	sqlQuery += where

	return sqlQuery
}

func QueryUpdateRole(role Role) string {
	sqlQuery := "update roles set Name = '" + role.Name + "' where Id = ? returning *"

	return sqlQuery
}

func QueryUpdateRoleIsActive(isActive bool) string {
	sqlQuery := "update roles set Is_Active = '" + strconv.FormatBool(isActive) + "' where Id = ?"

	return sqlQuery
}