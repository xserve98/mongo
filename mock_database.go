package mongo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	mgo_v2 "gopkg.in/mgo.v2"
)

// MockDatabaser is a mock of Databaser interface
type MockDatabaser struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaserMockRecorder
}

// MockDatabaserMockRecorder is the mock recorder for mockDatabaser
type MockDatabaserMockRecorder struct {
	mock *MockDatabaser
}

// newMockDatabaser creates a new mock instance
func newMockDatabaser(ctrl *gomock.Controller) *MockDatabaser {
	mock := &MockDatabaser{ctrl: ctrl}
	mock.recorder = &MockDatabaserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDatabaser) EXPECT() *MockDatabaserMockRecorder {
	return m.recorder
}

// AddUser mocks base method
func (m *MockDatabaser) AddUser(arg0, arg1 string, arg2 bool) error {
	ret := m.ctrl.Call(m, "AddUser", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddUser indicates an expected call of AddUser
func (mr *MockDatabaserMockRecorder) AddUser(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddUser", reflect.TypeOf((*MockDatabaser)(nil).AddUser), arg0, arg1, arg2)
}

// C mocks base method
func (m *MockDatabaser) C(arg0 string) Collectioner {
	ret := m.ctrl.Call(m, "C", arg0)
	ret0, _ := ret[0].(Collectioner)
	return ret0
}

// C indicates an expected call of C
func (mr *MockDatabaserMockRecorder) C(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "C", reflect.TypeOf((*MockDatabaser)(nil).C), arg0)
}

// CollectionNames mocks base method
func (m *MockDatabaser) CollectionNames() ([]string, error) {
	ret := m.ctrl.Call(m, "CollectionNames")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectionNames indicates an expected call of CollectionNames
func (mr *MockDatabaserMockRecorder) CollectionNames() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectionNames", reflect.TypeOf((*MockDatabaser)(nil).CollectionNames))
}

// DropDatabase mocks base method
func (m *MockDatabaser) DropDatabase() error {
	ret := m.ctrl.Call(m, "DropDatabase")
	ret0, _ := ret[0].(error)
	return ret0
}

// DropDatabase indicates an expected call of DropDatabase
func (mr *MockDatabaserMockRecorder) DropDatabase() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DropDatabase", reflect.TypeOf((*MockDatabaser)(nil).DropDatabase))
}

// FindRef mocks base method
func (m *MockDatabaser) FindRef(arg0 *mgo_v2.DBRef) Querier {
	ret := m.ctrl.Call(m, "FindRef", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// FindRef indicates an expected call of FindRef
func (mr *MockDatabaserMockRecorder) FindRef(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindRef", reflect.TypeOf((*MockDatabaser)(nil).FindRef), arg0)
}

// GridFS mocks base method
func (m *MockDatabaser) GridFS(arg0 string) *mgo_v2.GridFS {
	ret := m.ctrl.Call(m, "GridFS", arg0)
	ret0, _ := ret[0].(*mgo_v2.GridFS)
	return ret0
}

// GridFS indicates an expected call of GridFS
func (mr *MockDatabaserMockRecorder) GridFS(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GridFS", reflect.TypeOf((*MockDatabaser)(nil).GridFS), arg0)
}

// Login mocks base method
func (m *MockDatabaser) Login(arg0, arg1 string) error {
	ret := m.ctrl.Call(m, "Login", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Login indicates an expected call of Login
func (mr *MockDatabaserMockRecorder) Login(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockDatabaser)(nil).Login), arg0, arg1)
}

// Logout mocks base method
func (m *MockDatabaser) Logout() {
	m.ctrl.Call(m, "Logout")
}

// Logout indicates an expected call of Logout
func (mr *MockDatabaserMockRecorder) Logout() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Logout", reflect.TypeOf((*MockDatabaser)(nil).Logout))
}

// RemoveUser mocks base method
func (m *MockDatabaser) RemoveUser(arg0 string) error {
	ret := m.ctrl.Call(m, "RemoveUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveUser indicates an expected call of RemoveUser
func (mr *MockDatabaserMockRecorder) RemoveUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveUser", reflect.TypeOf((*MockDatabaser)(nil).RemoveUser), arg0)
}

// Run mocks base method
func (m *MockDatabaser) Run(arg0, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Run", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run
func (mr *MockDatabaserMockRecorder) Run(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockDatabaser)(nil).Run), arg0, arg1)
}

// UpsertUser mocks base method
func (m *MockDatabaser) UpsertUser(arg0 *mgo_v2.User) error {
	ret := m.ctrl.Call(m, "UpsertUser", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertUser indicates an expected call of UpsertUser
func (mr *MockDatabaserMockRecorder) UpsertUser(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertUser", reflect.TypeOf((*MockDatabaser)(nil).UpsertUser), arg0)
}

// With mocks base method
func (m *MockDatabaser) With(arg0 *mgo_v2.Session) Databaser {
	ret := m.ctrl.Call(m, "With", arg0)
	ret0, _ := ret[0].(Databaser)
	return ret0
}

// With indicates an expected call of With
func (mr *MockDatabaserMockRecorder) With(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "With", reflect.TypeOf((*MockDatabaser)(nil).With), arg0)
}
