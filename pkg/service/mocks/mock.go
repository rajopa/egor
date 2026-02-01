package mock_service

import (
	"reflect"

	domain "github.com/egor/watcher/pkg/model"
	"go.uber.org/mock/gomock"
)

type MockAuthorization struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationMockRecorder
}

type MockAuthorizationMockRecorder struct {
	mock *MockAuthorization
}

func NewMockAuthorization(ctrl *gomock.Controller) *MockAuthorization {
	mock := &MockAuthorization{ctrl: ctrl}
	mock.recorder = &MockAuthorizationMockRecorder{mock}
	return mock
}

func (m *MockAuthorization) EXPECT() *MockAuthorizationMockRecorder {
	return m.recorder
}

func (m *MockAuthorization) CreateUser(user domain.User) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockAuthorization)(nil).CreateUser), user)
}

func (m *MockAuthorization) GenerateToken(username, password string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateToken", username, password)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationMockRecorder) GenerateToken(username, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateToken", reflect.TypeOf((*MockAuthorization)(nil).GenerateToken), username, password)
}

func (m *MockAuthorization) ParseToken(token string) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ParseToken", token)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockAuthorizationMockRecorder) ParseToken(token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ParseToken", reflect.TypeOf((*MockAuthorization)(nil).ParseToken), token)
}

// --- MockDomainTarget ---

type MockDomainTarget struct {
	ctrl     *gomock.Controller
	recorder *MockDomainTargetMockRecorder
}

type MockDomainTargetMockRecorder struct {
	mock *MockDomainTarget
}

func NewMockDomainTarget(ctrl *gomock.Controller) *MockDomainTarget {
	mock := &MockDomainTarget{ctrl: ctrl}
	mock.recorder = &MockDomainTargetMockRecorder{mock}
	return mock
}

func (m *MockDomainTarget) EXPECT() *MockDomainTargetMockRecorder {
	return m.recorder
}

func (m *MockDomainTarget) Create(userId int, target domain.Target) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", userId, target)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDomainTargetMockRecorder) Create(userId, target interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockDomainTarget)(nil).Create), userId, target)
}

func (m *MockDomainTarget) GetAll(userId int) ([]domain.Target, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", userId)
	ret0, _ := ret[0].([]domain.Target)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDomainTargetMockRecorder) GetAll(userId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDomainTarget)(nil).GetAll), userId)
}

func (m *MockDomainTarget) GetById(userId, targetId int) (domain.Target, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", userId, targetId)
	ret0, _ := ret[0].(domain.Target)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockDomainTargetMockRecorder) GetById(userId, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockDomainTarget)(nil).GetById), userId, targetId)
}

func (m *MockDomainTarget) Delete(userId, targetId int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", userId, targetId)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDomainTargetMockRecorder) Delete(userId, targetId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDomainTarget)(nil).Delete), userId, targetId)
}

func (m *MockDomainTarget) Update(userId, targetId int, input domain.UpdateTargetInput) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", userId, targetId, input)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockDomainTargetMockRecorder) Update(userId, targetId, input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockDomainTarget)(nil).Update), userId, targetId, input)
}
