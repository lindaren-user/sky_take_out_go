package repository

import (
	"database/sql"
	"sky_take_out/model"
	"sky_take_out/utils"

	"go.uber.org/zap"
)

type EmployeeRepo interface {
	Login(username string, password string) (*model.Employee, error)
}

type employeeRepoImpl struct {
	conn *sql.DB
}

func NewEmployeeRepo(conn *sql.DB) EmployeeRepo {
	return &employeeRepoImpl{conn: conn}
}

func (e *employeeRepoImpl) Login(username string, password string) (*model.Employee, error) {
	query := "select id, name, username, password from employee where username = ? and password = ?" // mysql 使用 ? 占位

	employee := &model.Employee{} // var employee *model.Employee, 它是 nil，你不能对它的字段进行 Scan。会 panic。
	err := e.conn.QueryRow(query, username, password).Scan(&employee.Id, &employee.Name, &employee.Username, &employee.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			utils.Logger.Warn("用户名或密码错误")
			return nil, err
		}
		utils.Logger.Error("登录查询错误", zap.Error(err))
		return nil, err
	}
	return employee, nil
}
