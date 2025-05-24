package dto

type EmployeeLoginDTO struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type EmployeeSaveReqDTO struct {
	Id       int    `json:"id"`
	IdName   string `json:"idName"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Sex      string `json:"sex"`
	Username string `json:"username"`
}
