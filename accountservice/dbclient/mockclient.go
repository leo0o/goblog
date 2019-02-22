package dbclient

import (
	"github.com/stretchr/testify/mock"
	"github.com/leo0o/goblog/accountservice/model"
)

type MockBoltClient struct {
	mock.Mock
}

func (m *MockBoltClient)OpenBoltDb() {

}

func (m *MockBoltClient)Seed() {

}

func (m *MockBoltClient)QueryAccount(id string) (model.Account, error) {
	args := m.Mock.Called(id)
	return args.Get(0).(model.Account), args.Error(1)
}
