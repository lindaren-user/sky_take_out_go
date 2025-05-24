package service

// import (
// 	"sky_take_out/mock"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/stretchr/testify/assert"
// )

// func TestLogin_Success(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	mockRepo := mock.NewMockEmployeeRepo(ctrl)

// 	expectedEmployee := &model.Employee{
// 		Id:       1,
// 		Name:     "管理员",
// 		Username: "admin",
// 		Password: "123456",
// 	}

// 	mockRepo.EXPECT().GetUserByLogin("admin", "123456").Return(expectedEmployee, nil).Times(1)

// 	service := NewEmployeeService(mockRepo)

// 	employee, err := service.Login("admin", "123456")

// 	assert.NoError(t, err)
// 	assert.Equal(t, expectedEmployee, employee)
// }
