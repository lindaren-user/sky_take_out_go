package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/repository"
	"sky_take_out/result"
	service "sky_take_out/service/categoryService"
	"sky_take_out/utils"
	"strconv"
	"strings"
)

func CategoryMakeHandler(db *sql.DB) {
	repo := repository.NewCategoryRepo(db)
	category := service.NewCategoryService(repo)
	categoryController := NewCategoryController(category)

	http.Handle("/admin/category", utils.JWTAdminMiddleware(http.HandlerFunc(categoryController.Save)))
	http.Handle("/admin/category/delete", utils.JWTAdminMiddleware(http.HandlerFunc(categoryController.Delete)))
	http.Handle("/admin/category/page", utils.JWTAdminMiddleware(http.HandlerFunc(categoryController.Page)))
	http.Handle("/admin/category/update", utils.JWTAdminMiddleware(http.HandlerFunc(categoryController.Update)))
	http.Handle("/admin/category/list", utils.JWTAdminMiddleware(http.HandlerFunc(categoryController.GetBySort)))
	http.Handle("/admin/category/status/", utils.JWTAdminMiddleware(http.HandlerFunc(categoryController.StartAndStop)))
}

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(service *service.CategoryService) *CategoryController {
	return &CategoryController{service: service}
}

func (c *CategoryController) Save(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		result.Error(w, "请求方式错误")
		return
	}

	var category *model.Category
	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		result.Error(w, "请求体解析失败")
		return
	}

	if err := c.service.Save(r.Context(), category); err != nil {
		result.Error(w, "新增失败")
		return
	}

	result.Success(w, "新增成功", nil)
}

func (c *CategoryController) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		result.Error(w, "请求方式错误")
		return
	}

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		result.Error(w, "id 参数错误")
		return
	}

	if err := c.service.Delete(id); err != nil {
		result.Error(w, "删除失败")
		return
	}

	result.Success(w, "删除成功", nil)
}

func (c *CategoryController) Page(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		result.Error(w, "请求方式错误")
		return
	}

	name := r.URL.Query().Get("name")
	pageStr := r.URL.Query().Get("page")
	pageSizeStr := r.URL.Query().Get("pageSize")
	typeStr := r.URL.Query().Get("type")

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
	ctype, err := strconv.Atoi(typeStr)
	if err != nil {
		result.Error(w, "ctype 参数错误")
	}

	employeePageRespDtO, err := c.service.Page(name, page, pageSize, ctype)
	if err != nil {
		result.Error(w, "查询失败")
		return
	}

	result.Success(w, "查询成功", employeePageRespDtO)
}

func (c *CategoryController) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		result.Error(w, "请求方式错误")
		return
	}

	var categoryUpdateDTO dto.CategoryUpdateDTO
	if err := json.NewDecoder(r.Body).Decode(&categoryUpdateDTO); err != nil {
		result.Error(w, "请求体解析失败")
		return
	}

	if err := c.service.UpdateInfo(r.Context(), categoryUpdateDTO); err != nil {
		result.Error(w, "更新失败")
		return
	}
	result.Success(w, "更新成功", nil)
}

func (c *CategoryController) GetBySort(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		result.Error(w, "请求方式错误")
		return
	}

	typeStr := r.URL.Query().Get("sort")
	ctype, err := strconv.Atoi(typeStr)
	if err != nil {
		result.Error(w, "sort 参数错误")
		return
	}

	category, err := c.service.GetByType(ctype)
	if err != nil {
		result.Error(w, "查询错误")
		return
	}
	result.Success(w, "查询成功", category)
}

func (c *CategoryController) StartAndStop(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		result.Error(w, "请求方式错误")
		return
	}

	path := r.URL.Path
	parts := strings.Split(strings.Trim(path, "/"), "/")
	if len(parts) < 4 {
		result.Error(w, "路径出错")
		return
	}
	statusStr := parts[len(parts)-1]

	idStr := r.URL.Query().Get("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		result.Error(w, "id 参数错误")
		return
	}

	status, err := strconv.Atoi(statusStr)
	if err != nil {
		result.Error(w, "status 参数错误")
		return
	}

	if err := c.service.StartAndStop(r.Context(), id, status); err != nil {
		result.Error(w, "执行失败")
		return
	}
	result.Success(w, "执行失败", nil)
}
