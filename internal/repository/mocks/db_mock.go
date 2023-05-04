// Code generated by MockGen. DO NOT EDIT.
// Source: internal/ports/repository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	models "github.com/michaelgbenle/ZeinaMFI/internal/models"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateUser mocks base method.
func (m *MockRepository) CreateUser(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateUser indicates an expected call of CreateUser.
func (mr *MockRepositoryMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockRepository)(nil).CreateUser), user)
}

// Deposit mocks base method.
func (m *MockRepository) Deposit(money *models.Money, creditor *models.User) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Deposit", money, creditor)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Deposit indicates an expected call of Deposit.
func (mr *MockRepositoryMockRecorder) Deposit(money, creditor interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Deposit", reflect.TypeOf((*MockRepository)(nil).Deposit), money, creditor)
}

// FindUserByAccountNos mocks base method.
func (m *MockRepository) FindUserByAccountNos(account string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByAccountNos", account)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByAccountNos indicates an expected call of FindUserByAccountNos.
func (mr *MockRepositoryMockRecorder) FindUserByAccountNos(account interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByAccountNos", reflect.TypeOf((*MockRepository)(nil).FindUserByAccountNos), account)
}

// FindUserByEmail mocks base method.
func (m *MockRepository) FindUserByEmail(email string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserByEmail", email)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserByEmail indicates an expected call of FindUserByEmail.
func (mr *MockRepositoryMockRecorder) FindUserByEmail(email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserByEmail", reflect.TypeOf((*MockRepository)(nil).FindUserByEmail), email)
}

// FindUserById mocks base method.
func (m *MockRepository) FindUserById(Id string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindUserById", Id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindUserById indicates an expected call of FindUserById.
func (mr *MockRepositoryMockRecorder) FindUserById(Id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindUserById", reflect.TypeOf((*MockRepository)(nil).FindUserById), Id)
}

// GetAllUsers mocks base method.
func (m *MockRepository) GetAllUsers() (*[]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].(*[]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockRepositoryMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockRepository)(nil).GetAllUsers))
}

// GetTransactions mocks base method.
func (m *MockRepository) GetTransactions(accountNo string) (*[]models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTransactions", accountNo)
	ret0, _ := ret[0].(*[]models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTransactions indicates an expected call of GetTransactions.
func (mr *MockRepositoryMockRecorder) GetTransactions(accountNo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTransactions", reflect.TypeOf((*MockRepository)(nil).GetTransactions), accountNo)
}

// LockSavings mocks base method.
func (m *MockRepository) LockSavings(money *models.Money, user *models.User) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LockSavings", money, user)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LockSavings indicates an expected call of LockSavings.
func (mr *MockRepositoryMockRecorder) LockSavings(money, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LockSavings", reflect.TypeOf((*MockRepository)(nil).LockSavings), money, user)
}

// TokenInBlacklist mocks base method.
func (m *MockRepository) TokenInBlacklist(token *string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TokenInBlacklist", token)
	ret0, _ := ret[0].(bool)
	return ret0
}

// TokenInBlacklist indicates an expected call of TokenInBlacklist.
func (mr *MockRepositoryMockRecorder) TokenInBlacklist(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TokenInBlacklist", reflect.TypeOf((*MockRepository)(nil).TokenInBlacklist), token)
}

// Withdraw mocks base method.
func (m *MockRepository) Withdraw(money *models.Money, user *models.User) (*models.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", money, user)
	ret0, _ := ret[0].(*models.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockRepositoryMockRecorder) Withdraw(money, user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockRepository)(nil).Withdraw), money, user)
}
