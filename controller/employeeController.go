package controller

import (
	"encoding/json"
	"net/http"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/result"
	service "sky_take_out/service/employeeService"
	"sky_take_out/utils"

	"go.uber.org/zap"
)

type EmployeeController struct {
	service *service.EmployeeService
}

func NewEmployeeController(service *service.EmployeeService) *EmployeeController {
	return &EmployeeController{service: service}
}

func (e *EmployeeController) Login(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		result.Error(w, "请求方式错误")
		return
	}

	var em dto.EmployeeLoginDTO
	if err := json.NewDecoder(r.Body).Decode(&em); err != nil {
		utils.Logger.Error("解析请求体失败", zap.Error(err))
		result.Error(w, "解析请求体失败")
		return
	}

	employee, err := e.service.Login(em.Username, em.Password)
	if err != nil {
		result.Error(w, "登录失败")
		return
	}

	token, err := utils.GenerateJWT(employee.Id)
	if err != nil {
		result.Error(w, "token生成失败")
		return
	}

	result.Success(w, "登录成功", &struct {
		Id       int    `json:"id"`
		Name     string `json:"name"`
		Username string `json:"userName"`
		Token    string `json:"token"`
	}{
		Id:       employee.Id,
		Name:     employee.Name,
		Username: employee.Username,
		Token:    token,
	})
}

func (e *EmployeeController) Logout(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		result.Error(w, "请求方式错误")
		return
	}

	result.Success(w, "登出成功", nil)
}

func (e *EmployeeController) Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		result.Error(w, "请求方式错误")
		return
	}

	var employee *model.Employee
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		result.Error(w, "解析请求体失败")
		return
	}

	if err := e.service.Save(employee); err != nil {
		result.Error(w, "新增失败")
		return
	}
	result.Success(w, "新增成功", nil)
}
