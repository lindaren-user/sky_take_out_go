package model

import "time"

type Category struct {
	Id         int       `json:"id"`
	Type       int       `json:"type"`
	Name       string    `json:"name"`
	Sort       int       `json:"sort"`
	Status     int       `json:"status"`
	CreateTime time.Time `json:"create_time"`
	CreateUser int       `json:"create_user"`
	UpdateTime time.Time `json:"update_time"`
	UpdateUser int       `json:"update_user"`
}
