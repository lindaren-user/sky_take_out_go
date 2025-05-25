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

	GetUserByPage(name string, page int, pageSize int) (total int, employees []*model.Employee, err error)

	StartAndStop(employeeId int, status int) error

	GetInfo(id int) (*model.Employee, error)

	UpdateInfo(employeeUpdateReqDTO *dto.EmployeeUpdateReqDTO) error
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

func (e *employeeRepoImpl) GetUserByPage(name string, page int, pageSize int) (total int, employees []*model.Employee, err error) {
	// 查询总数
	countQuery := "select COUNT(*) from employee where name like concat('%', ?, '%')"
	err = e.conn.QueryRow(countQuery, name).Scan(&total)
	if err != nil {
		utils.Logger.Error("数据库查询出错", zap.Error(err))
		return
	}

	// 分页查询
	offset := (page - 1) * pageSize
	query := "select id, name, username, password, phone, sex, id_number, status, create_time, update_time, create_user, update_user from employee where name like concat('%', ?, '%') limit ? offset ?"
	rows, err := e.conn.Query(query, name, pageSize, offset)
	if err != nil {
		utils.Logger.Error("数据库查询失败", zap.Error(err))
		return
	}
	defer rows.Close()

	employees = make([]*model.Employee, 0) // 序列化为 []，不会是 null

	for rows.Next() {
		emp := &model.Employee{}
		err = rows.Scan(
			&emp.Id,
			&emp.Name,
			&emp.Username,
			&emp.Password,
			&emp.Phone,
			&emp.Sex,
			&emp.IdNumber,
			&emp.Status,
			&emp.CreateTime,
			&emp.UpdateTime,
			&emp.CreateUser,
			&emp.UpdateUser,
		)
		if err != nil {
			utils.Logger.Error("读取行失败", zap.Error(err))
			return
		}
		employees = append(employees, emp)
	}

	if err = rows.Err(); err != nil {
		utils.Logger.Error("遍历行失败", zap.Error(err))
		return
	}

	return
}

func (e *employeeRepoImpl) StartAndStop(employeeId int, status int) error {
	update := "update employee set status = ? where id = ?"

	if _, err := e.conn.Exec(update, status, employeeId); err != nil {
		utils.Logger.Error("更新状态失败", zap.Error(err))
		return err
	}

	return nil
}

func (e *employeeRepoImpl) GetInfo(id int) (*model.Employee, error) {
	query := "select id, name, username, password, phone, sex, id_number, status, create_time, update_time, create_user, update_user from employee where id = ?"

	employee := &model.Employee{}
	err := e.conn.QueryRow(query, id).Scan(
		&employee.Id,
		&employee.Name,
		&employee.Username,
		&employee.Password,
		&employee.Phone,
		&employee.Sex,
		&employee.IdNumber,
		&employee.Status,
		&employee.CreateTime,
		&employee.UpdateTime,
		&employee.CreateUser,
		&employee.UpdateUser,
	)
	if err != nil {
		utils.Logger.Error("查询失败", zap.Error(err))
		return nil, err
	}

	return employee, nil
}

func (e *employeeRepoImpl) UpdateInfo(employeeUpdateReqDTO *dto.EmployeeUpdateReqDTO) error {
	update := `
        update employee set 
            name = ?, 
            username = ?, 
            phone = ?, 
            sex = ?, 
            id_number = ?, 
            update_time = ?, 
            update_user = ?
        where id = ?
    `

	_, err := e.conn.Exec(
		update,
		employeeUpdateReqDTO.Name,
		employeeUpdateReqDTO.Username,
		employeeUpdateReqDTO.Phone,
		employeeUpdateReqDTO.Sex,
		employeeUpdateReqDTO.IdNumber,
		employeeUpdateReqDTO.UpdateTime,
		employeeUpdateReqDTO.UpdateUser,
		employeeUpdateReqDTO.Id,
	)
	if err != nil {
		utils.Logger.Error("更新失败", zap.Error(err))
		return err
	}

	return nil
}
