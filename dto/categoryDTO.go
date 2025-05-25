package dto

import (
	"sky_take_out/model"
	"time"
)

type CategorySaveRepDTO struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Sort int    `json:"sort"`
	Type int    `json:"type"`
}

type CategoryPageRespDTO struct {
	Total   int              `json:"total"`
	Records []model.Category `json:"records"`
}

type CategoryUpdateDTO struct {
	CategorySaveRepDTO
	UpdateTime time.Time `json:"update_time"`
	UpdateUser int       `json:"update_user"`
}

type CategoryStatusDTO struct {
	Id         int       `json:"id"`
	Status     int       `json:"status"`
	UpdateTime time.Time `json:"update_time"`
	UpdateUser int       `json:"update_user"`
}
