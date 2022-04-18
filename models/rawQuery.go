package models

import (
	"fmt"
	"strconv"
)

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

func QueryDeleteDeletePermission() string {
	sql := "delete from role_permisions where role_id = ?"
	return sql
}

func QueryUpdateRolePermission(id_permission Permission, id_role uint) (string) {
	//sqlQuery := fmt.Sprintf("delete from role_permisions where role_id = %v ",id_role)
	sqlQuery := fmt.Sprintf("insert into role_permisions (role_id,permission_id) values(%v,%v)",id_role,id_permission.Id)
	return sqlQuery
}

func QueryUpdateProduct(product Product) string {
	sqlQuery := "update products set "
	if product.Title != "" {
		sqlQuery += "title = '"+ product.Title +"' "
	}
	if product.Description != "" {
		sqlQuery += "description = '"+product.Description+"' "
	}
	if product.Image != "" {
		sqlQuery += "image = '"+product.Image+"' "
	}
	if product.Price != 0 {
		sqlQuery += fmt.Sprintf("price = '%v' ", product.Price)
	}

	where := "where id = ? returning *"

	sqlQuery += where
	return sqlQuery
}