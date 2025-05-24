package repository

import (
	"database/sql"
	"sky_take_out/dto"
	"sky_take_out/model"
	"sky_take_out/utils"

	"go.uber.org/zap"
)

//go:generate mockgen -source=employeeRepo.go -destination=../mock/employeeRepoMock.go -package=mock
type EmployeeRepo interface {
	GetUserByLogin(username string, password string) (*dto.EmployeeLoginDTO, error)

	Insert(employee *model.Employee) error
}

type employeeRepoImpl struct {
	conn *sql.DB
}

func NewEmployeeRepo(conn *sql.DB) EmployeeRepo {
	return &employeeRepoImpl{conn: conn}
}

func (e *employeeRepoImpl) GetUserByLogin(username string, password string) (*dto.EmployeeLoginDTO, error) {
	query := "select id, name, username, password from employee where username = ? and password = ?" // mysql 使用 ? 占位

	employee := &dto.EmployeeLoginDTO{} // var employee *model.Employee, 它是 nil，你不能对它的字段进行 Scan。会 panic。
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

func (e *employeeRepoImpl) Insert(employee *model.Employee) error {
	insert := `
        INSERT INTO employee (
            name, username, password, phone, sex, id_number, status, create_time, update_time, create_user, update_user
        ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
    `
	_, err := e.conn.Exec(
		insert,
		employee.Name,
		employee.Username,
		employee.Password,
		employee.Phone,
		employee.Sex,
		employee.IdNumber,
		employee.Status,
		employee.CreateTime,
		employee.UpdateTime,
		employee.CreateUser,
		employee.UpdateUser,
	)
	if err != nil {
		utils.Logger.Error("插入员工失败", zap.Error(err))
		return err
	}
	return nil
}
