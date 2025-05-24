package service

import (
	"sky_take_out/constant"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/repository"
	"time"
)

type EmployeeService struct {
	repo repository.EmployeeRepo
}

func NewEmployeeService(repo repository.EmployeeRepo) *EmployeeService {
	return &EmployeeService{repo: repo}
}

func (e *EmployeeService) Login(username string, password string) (*dto.EmployeeLoginDTO, error) {
	return e.repo.GetUserByLogin(username, password)
}

func (e *EmployeeService) Save(employee *model.Employee) error {
	employee.Password = constant.Password
	employee.Status = constant.Status

	employee.CreateTime = time.Now()
	employee.UpdateTime = time.Now()

	// TODO 这是创建者的 ID
	employee.CreateUser = 1
	employee.UpdateUser = 1

	return e.repo.Insert(employee)
}
