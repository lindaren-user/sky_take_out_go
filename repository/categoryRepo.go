package repository

import (
	"database/sql"
	"go.uber.org/zap"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/utils"
)

type CategoryRepo interface {
	Insert(category *model.Category) error

	Delete(id int) error

	GetCategoriesByPage(name string, page int, pageSize int, ctype int) (total int, categories []model.Category, err error)

	UpdateInfo(categoryUpdateDTO dto.CategoryUpdateDTO) error

	GetByType(ctype int) (*model.Category, error)

	StartAndStop(statusDTO *dto.CategoryStatusDTO) error
}

type categoryRepoImpl struct {
	conn *sql.DB
}

func NewCategoryRepo(conn *sql.DB) CategoryRepo {
	return &categoryRepoImpl{conn: conn}
}

func (c *categoryRepoImpl) Insert(category *model.Category) error {
	insert := "insert into category (id, type, name, sort, status, create_time, create_user, update_time, update_user) values(?,?,?,?,?,?,?,?,?) "

	if _, err := c.conn.Exec(insert, &category.Id, &category.Type, &category.Name, &category.Sort, &category.Status, &category.CreateTime, &category.CreateUser, &category.UpdateTime, &category.UpdateUser); err != nil {
		utils.Logger.Error("更新失败", zap.Error(err))
		return err
	}
	return nil
}

func (c *categoryRepoImpl) Delete(id int) error {
	deleteSql := "delete from category where id=?"
	if _, err := c.conn.Exec(deleteSql, id); err != nil {
		utils.Logger.Error("删除失败", zap.Error(err))
		return err
	}
	return nil
}

func (c *categoryRepoImpl) GetCategoriesByPage(name string, page int, pageSize int, ctype int) (total int, categories []model.Category, err error) {
	query := "select COUNT(*) from category"

	if err = c.conn.QueryRow(query).Scan(&total); err != nil {
		utils.Logger.Error("查询出错", zap.Error(err))
		return
	}

	offset := (page - 1) * pageSize
	keyword := "%" + name + "%" // name 是你要搜索的关键字

	query = `
    select id, type, name, sort, status, create_time, create_user, update_time, update_user 
    from category 
    where name like ? 
    limit ? offset ?
`

	rows, err := c.conn.Query(query, keyword, pageSize, offset)
	if err != nil {
		utils.Logger.Error("查询出错", zap.Error(err))
		return
	}
	defer rows.Close()

	categories = make([]model.Category, 0)

	for rows.Next() {
		category := model.Category{}
		err = rows.Scan(
			&category.Id,
			&category.Type,
			&category.Name,
			&category.Sort,
			&category.Status,
			&category.CreateTime,
			&category.CreateUser,
			&category.UpdateTime,
			&category.UpdateUser,
		)
		if err != nil {
			utils.Logger.Error("读取失败", zap.Error(err))
			return
		}

		categories = append(categories, category)
	}

	if rows.Err() != nil {
		utils.Logger.Error("遍历行出错", zap.Error(err))
		return
	}
	return
}

func (c *categoryRepoImpl) UpdateInfo(categoryUpdateDTO dto.CategoryUpdateDTO) error {
	update := "update category set name = ? and sort = ? and type = ? and update_time = ? and update_user = ? where id = ?"

	if _, err := c.conn.Exec(update, categoryUpdateDTO.Name, categoryUpdateDTO.Sort, categoryUpdateDTO.Type, categoryUpdateDTO.Id, categoryUpdateDTO.UpdateTime, categoryUpdateDTO.UpdateUser); err != nil {
		utils.Logger.Error("更新错误", zap.Error(err))
		return err
	}
	return nil
}

func (c *categoryRepoImpl) GetByType(ctype int) (*model.Category, error) {
	query := "SELECT id, type, name, sort, status, create_time, create_user, update_time, update_user FROM category WHERE type = ?"

	category := &model.Category{}
	err := c.conn.QueryRow(query, ctype).Scan(
		&category.Id,
		&category.Type,
		&category.Name,
		&category.Sort,
		&category.Status,
		&category.CreateTime,
		&category.CreateUser,
		&category.UpdateTime,
		&category.UpdateUser,
	)
	if err != nil {
		utils.Logger.Error("<UNK>", zap.Error(err))
		return nil, err
	}
	return category, nil
}

func (c *categoryRepoImpl) StartAndStop(statusDTO *dto.CategoryStatusDTO) error {
	update := "update category set status = ? and update_time = ? and update_user = ? where id = ?"

	if _, err := c.conn.Exec(update, statusDTO.Status, statusDTO.UpdateTime, statusDTO.UpdateUser, statusDTO.Id); err != nil {
		utils.Logger.Error("<UNK>", zap.Error(err))
		return err
	}
	return nil
}
