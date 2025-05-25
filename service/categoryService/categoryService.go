package service

import (
	"context"
	"fmt"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/repository"
	"sky_take_out/utils"
	"time"
)

type CategoryService struct {
	repo repository.CategoryRepo
}

func NewCategoryService(repo repository.CategoryRepo) *CategoryService {
	return &CategoryService{repo: repo}
}

func (c *CategoryService) Save(ctx context.Context, category *model.Category) error {
	adminID, ok := utils.GetAdminID(ctx)
	if !ok {
		utils.Logger.Error("获取 admin_id 错误")
		return fmt.Errorf("获取 admin_id 错误")
	}

	category.CreateTime = time.Now()
	category.UpdateTime = time.Now()

	category.CreateUser = adminID
	category.UpdateUser = adminID

	return c.repo.Insert(category)
}

func (c *CategoryService) Delete(id int) error {
	return c.repo.Delete(id)
}

func (c *CategoryService) Page(name string, page int, pageSize int, ctype int) (*dto.CategoryPageRespDTO, error) {
	total, categories, err := c.repo.GetCategoriesByPage(name, page, pageSize, ctype)
	if err != nil {
		return nil, err
	}

	categoryPageRespDTO := &dto.CategoryPageRespDTO{
		Total:   total,
		Records: categories,
	}
	return categoryPageRespDTO, nil
}

func (c *CategoryService) UpdateInfo(ctx context.Context, categoryUpdateDTO dto.CategoryUpdateDTO) error {
	adminID, ok := utils.GetAdminID(ctx)
	if !ok {
		utils.Logger.Error("获取 admin_id 错误")
		return fmt.Errorf("获取 admin_id 错误")
	}

	categoryUpdateDTO.UpdateTime = time.Now()
	categoryUpdateDTO.UpdateUser = adminID

	return c.repo.UpdateInfo(categoryUpdateDTO)
}

func (c *CategoryService) GetByType(ctype int) (category *model.Category, err error) {
	return c.repo.GetByType(ctype)
}

func (c *CategoryService) StartAndStop(ctx context.Context, status int, id int) error {
	adminID, ok := utils.GetAdminID(ctx)
	if !ok {
		utils.Logger.Error("获取 admin_id 错误")
		return fmt.Errorf("获取 admin_id 错误")
	}

	statusDTO := &dto.CategoryStatusDTO{}
	statusDTO.Status = status
	statusDTO.Id = id
	statusDTO.UpdateTime = time.Now()
	statusDTO.UpdateUser = adminID

	return c.repo.StartAndStop(statusDTO)
}
