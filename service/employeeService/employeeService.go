package service

import (
	"context"
	"fmt"
	"sky_take_out/constant"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/repository"
	"sky_take_out/utils"
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

func (e *EmployeeService) Save(ctx context.Context, employee *model.Employee) error {
	employee.Password = constant.Password
	employee.Status = constant.Status

	employee.CreateTime = time.Now()
	employee.UpdateTime = time.Now()

	adminID, ok := utils.GetAdminID(ctx)
	if !ok {
		utils.Logger.Error("获取 admin_id 错误")
		return fmt.Errorf("获取 admin_id 错误")
	}

	employee.CreateUser = adminID
	employee.UpdateUser = adminID

	return e.repo.Insert(employee)
}

func (e *EmployeeService) Page(name string, page int, pageSize int) (*dto.EmployeePageRespDTO, error) {
	total, employees, err := e.repo.GetUserByPage(name, page, pageSize)
	if err != nil {
		return nil, err
	}

	employeePageRespDtO := &dto.EmployeePageRespDTO{
		Total:   total,
		Records: employees,
	}
	return employeePageRespDtO, nil
}
