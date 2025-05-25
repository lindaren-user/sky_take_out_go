package dto

import (
	"sky_take_out/model"
	"time"
)

type EmployeeLoginDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmployeeSaveReqDTO struct {
	Id       int    `json:"id"`
	IdNumber string `json:"id_number"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Username string `json:"username"`
}

type EmployeePageRespDTO struct {
	Total   int               `json:"total"`
	Records []*model.Employee `json:"records"` // TODO：[]model.Employee 的区别
}

type EmployeeUpdateReqDTO struct {
	EmployeeSaveReqDTO
	UpdateTime time.Time `json:"update_time"`
	UpdateUser int       `json:"update_user"`
}

type EmployeeStatusDTO struct {
	Id         int       `json:"id"`
	Status     int       `json:"status"`
	UpdateTime time.Time `json:"update_time"`
	UpdateUser int       `json:"update_user"`
}
