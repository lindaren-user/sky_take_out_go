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

// TODO：禁用的账号无法登录，密码加密
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

func (e *EmployeeService) StartAndStop(ctx context.Context, id int, status int) error {
	adminID, ok := utils.GetAdminID(ctx)
	if !ok {
		utils.Logger.Error("获取 admin_id 错误")
		return fmt.Errorf("获取 admin_id 错误")
	}

	statusDTO := &dto.EmployeeStatusDTO{}

	statusDTO.Status = status
	statusDTO.Id = id
	statusDTO.UpdateTime = time.Now() // TODO： update_time DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
	statusDTO.UpdateUser = adminID

	return e.repo.StartAndStop(statusDTO)
}

func (e *EmployeeService) GetInfo(id int) (*model.Employee, error) {
	return e.repo.GetInfo(id)
}

func (e *EmployeeService) UpdateInfo(ctx context.Context, employeeUpdateReqDTO *dto.EmployeeUpdateReqDTO) error {
	adminID, ok := utils.GetAdminID(ctx)
	if !ok {
		utils.Logger.Error("获取 admin_id 错误")
		return fmt.Errorf("获取 admin_id 错误")
	}

	employeeUpdateReqDTO.UpdateTime = time.Now()
	employeeUpdateReqDTO.UpdateUser = adminID

	return e.repo.UpdateInfo(employeeUpdateReqDTO)
}
