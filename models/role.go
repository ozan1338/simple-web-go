package models

type Role struct {
	Id       uint   `json:"id"`
	Name     string `json:"name"`
	IsActive bool   `json:"is_active"`
}