package model

import "time"

type Employee struct {
	Id         int       `json:"id"`
	Name       string    `json:"name"`
	Username   string    `json:"username"`
	Password   string    `json:"password"`
	Phone      string    `json:"phone"`
	Sex        string    `json:"sex"`
	IdNumber   string    `json:"idNumber"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"createTime"`
	UpdateTime time.Time `json:"updateTime"`
	CreateUser int       `json:"createUser"`
	UpdateUser int       `json:"updateUser"`
}
