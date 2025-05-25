package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sky_take_out/database"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/repository"
	"sky_take_out/result"
	service "sky_take_out/service/employeeService"
	"sky_take_out/utils"
	"strconv"
	"strings"

	"go.uber.org/zap"
)

func EmployeeMakeHandler(db *sql.DB) {
	repo := repository.NewEmployeeRepo(database.GetDB())
	employeeService := service.NewEmployeeService(repo)
	employeeController := NewEmployeeController(employeeService)

	http.HandleFunc("/admin/employee/login", employeeController.Login)
	http.HandleFunc("/admin/employee/logout", employeeController.Logout)
	http.Handle("/admin/employee", utils.JWTAdminMiddleware(http.HandlerFunc(employeeController.Save)))
	http.Handle("/admin/employee/page", utils.JWTAdminMiddleware(http.HandlerFunc(employeeController.Page)))
	http.Handle("/admin/employee/status/", utils.JWTAdminMiddleware(http.HandlerFunc(employeeController.StartAndStop)))
	http.Handle("/admin/employee/", utils.JWTAdminMiddleware(http.HandlerFunc(employeeController.GetInfo)))
	// http.Handle("/admin/employee", utils.JWTAdminMiddleware(http.HandlerFunc(employeeController.UpdateInfo))) // go 原生不允许相同的路由，gin 可以
	http.Handle("/admin/employee/update", utils.JWTAdminMiddleware(http.HandlerFunc(employeeController.UpdateInfo)))
}

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
		result.Error(w, "token 生成失败")
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

	if err := e.service.Save(r.Context(), employee); err != nil {
		result.Error(w, "新增失败")
		return
	}
	result.Success(w, "新增成功", nil)
}

func (e *EmployeeController) Page(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		result.Error(w, "请求方式错误")
		return
	}

	name := r.URL.Query().Get("name")
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")

	// TODO：怎么处理这些参数才是最优解
	page, err := strconv.Atoi(pageStr)
	if err != nil {
		result.Error(w, "page 参数错误")
		return
	}
	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil {
		result.Error(w, "pageSize 参数错误")
		return
	}

	employeePageRespDtO, err := e.service.Page(name, page, pageSize)
	if err != nil {
		result.Error(w, "查询失败")
		return
	}

	result.Success(w, "查询成功", employeePageRespDtO)
}

func (e *EmployeeController) StartAndStop(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		result.Error(w, "请求方式错误")
		return
	}

	// TODO：怎么优化
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 4 { // /admin/employee/status/{status}
		result.Error(w, "路径错误")
		return
	}
	statusStr := parts[len(parts)-1]

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		result.Error(w, "status 参数错误")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		result.Error(w, "id 参数错误")
		return
	}

	if err = e.service.StartAndStop(r.Context(), id, status); err != nil {
		result.Error(w, "启用/禁用出错")
		return
	}

	result.Success(w, "执行成功", nil)
}

func (e *EmployeeController) GetInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		result.Error(w, "请求方式错误")
		return
	}

	// TODO：怎么优化
	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 3 { // /admin/employee/{id}
		result.Error(w, "路径错误")
		return
	}
	idStr := parts[len(parts)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		result.Error(w, "id 参数错误")
		return
	}

	employee, err := e.service.GetInfo(id)
	if err != nil {
		result.Error(w, "查询失败")
		return
	}

	result.Success(w, "查询成功", employee)
}

func (e *EmployeeController) UpdateInfo(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		result.Error(w, "请求方式错误")
		return
	}

	var employeeUpdateReqDTO *dto.EmployeeUpdateReqDTO
	if err := json.NewDecoder(r.Body).Decode(&employeeUpdateReqDTO); err != nil {
		result.Error(w, "请求体解析失败")
		return
	}

	if err := e.service.UpdateInfo(r.Context(), employeeUpdateReqDTO); err != nil {
		result.Error(w, "更新失败")
		return
	}

	result.Success(w, "更新成功", nil)
}
