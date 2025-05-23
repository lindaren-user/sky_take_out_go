package controller

import (
	"encoding/json"
	"net/http"
	"sky_take_out/model"
	"sky_take_out/service"
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
		model.Error(w, "请求方式错误")
		return
	}

	var em model.Employee
	if err := json.NewDecoder(r.Body).Decode(&em); err != nil {
		utils.Logger.Error("解析请求体失败", zap.Error(err))
		model.Error(w, "解析请求体失败")
		return
	}

	employee, err := e.service.Login(em.Username, em.Password)
	if err != nil {
		model.Error(w, "登录失败")
		return
	}

	token, err := utils.GenerateJWT(employee.Id)
	if err != nil {
		model.Error(w, "token生成失败")
		return
	}

	model.Success(w, "登录成功", &struct {
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
