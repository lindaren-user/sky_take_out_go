package service

import (
	"sky_take_out/model"
	"sky_take_out/repository"
)

type EmployeeService struct {
	repo repository.EmployeeRepo
}

func NewEmployeeService(repo repository.EmployeeRepo) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (e *EmployeeService) Login(username string, password string) (*model.Employee, error) {
	return e.repo.Login(username, password)
}
