package mongo

import (
	reflect "reflect"
	time "time"

	gomock "github.com/golang/mock/gomock"
	mgo_v2 "gopkg.in/mgo.v2"
)

// MockQuerier is a mock of Querier interface
type MockQuerier struct {
	ctrl     *gomock.Controller
	recorder *MockQuerierMockRecorder
}

// MockQuerierMockRecorder is the mock recorder for mockQuerier
type MockQuerierMockRecorder struct {
	mock *MockQuerier
}

// newMockQuerier creates a new mock instance
func newMockQuerier(ctrl *gomock.Controller) *MockQuerier {
	mock := &MockQuerier{ctrl: ctrl}
	mock.recorder = &MockQuerierMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockQuerier) EXPECT() *MockQuerierMockRecorder {
	return m.recorder
}

// All mocks base method
func (m *MockQuerier) All(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "All", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// All indicates an expected call of All
func (mr *MockQuerierMockRecorder) All(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockQuerier)(nil).All), arg0)
}

// Apply mocks base method
func (m *MockQuerier) Apply(arg0 mgo_v2.Change, arg1 interface{}) (*mgo_v2.ChangeInfo, error) {
	ret := m.ctrl.Call(m, "Apply", arg0, arg1)
	ret0, _ := ret[0].(*mgo_v2.ChangeInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Apply indicates an expected call of Apply
func (mr *MockQuerierMockRecorder) Apply(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Apply", reflect.TypeOf((*MockQuerier)(nil).Apply), arg0, arg1)
}

// Batch mocks base method
func (m *MockQuerier) Batch(arg0 int) Querier {
	ret := m.ctrl.Call(m, "Batch", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Batch indicates an expected call of Batch
func (mr *MockQuerierMockRecorder) Batch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Batch", reflect.TypeOf((*MockQuerier)(nil).Batch), arg0)
}

// Comment mocks base method
func (m *MockQuerier) Comment(arg0 string) Querier {
	ret := m.ctrl.Call(m, "Comment", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Comment indicates an expected call of Comment
func (mr *MockQuerierMockRecorder) Comment(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Comment", reflect.TypeOf((*MockQuerier)(nil).Comment), arg0)
}

// Count mocks base method
func (m *MockQuerier) Count() (int, error) {
	ret := m.ctrl.Call(m, "Count")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count
func (mr *MockQuerierMockRecorder) Count() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockQuerier)(nil).Count))
}

// Distinct mocks base method
func (m *MockQuerier) Distinct(arg0 string, arg1 interface{}) error {
	ret := m.ctrl.Call(m, "Distinct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Distinct indicates an expected call of Distinct
func (mr *MockQuerierMockRecorder) Distinct(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Distinct", reflect.TypeOf((*MockQuerier)(nil).Distinct), arg0, arg1)
}

// Explain mocks base method
func (m *MockQuerier) Explain(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "Explain", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Explain indicates an expected call of Explain
func (mr *MockQuerierMockRecorder) Explain(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Explain", reflect.TypeOf((*MockQuerier)(nil).Explain), arg0)
}

// For mocks base method
func (m *MockQuerier) For(arg0 interface{}, arg1 func() error) error {
	ret := m.ctrl.Call(m, "For", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// For indicates an expected call of For
func (mr *MockQuerierMockRecorder) For(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "For", reflect.TypeOf((*MockQuerier)(nil).For), arg0, arg1)
}

// Hint mocks base method
func (m *MockQuerier) Hint(arg0 ...string) Querier {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Hint", varargs...)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Hint indicates an expected call of Hint
func (mr *MockQuerierMockRecorder) Hint(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hint", reflect.TypeOf((*MockQuerier)(nil).Hint), arg0...)
}

// Iter mocks base method
func (m *MockQuerier) Iter() *mgo_v2.Iter {
	ret := m.ctrl.Call(m, "Iter")
	ret0, _ := ret[0].(*mgo_v2.Iter)
	return ret0
}

// Iter indicates an expected call of Iter
func (mr *MockQuerierMockRecorder) Iter() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iter", reflect.TypeOf((*MockQuerier)(nil).Iter))
}

// Limit mocks base method
func (m *MockQuerier) Limit(arg0 int) Querier {
	ret := m.ctrl.Call(m, "Limit", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Limit indicates an expected call of Limit
func (mr *MockQuerierMockRecorder) Limit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Limit", reflect.TypeOf((*MockQuerier)(nil).Limit), arg0)
}

// LogReplay mocks base method
func (m *MockQuerier) LogReplay() Querier {
	ret := m.ctrl.Call(m, "LogReplay")
	ret0, _ := ret[0].(Querier)
	return ret0
}

// LogReplay indicates an expected call of LogReplay
func (mr *MockQuerierMockRecorder) LogReplay() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LogReplay", reflect.TypeOf((*MockQuerier)(nil).LogReplay))
}

// MapReduce mocks base method
func (m *MockQuerier) MapReduce(arg0 *mgo_v2.MapReduce, arg1 interface{}) (*mgo_v2.MapReduceInfo, error) {
	ret := m.ctrl.Call(m, "MapReduce", arg0, arg1)
	ret0, _ := ret[0].(*mgo_v2.MapReduceInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MapReduce indicates an expected call of MapReduce
func (mr *MockQuerierMockRecorder) MapReduce(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MapReduce", reflect.TypeOf((*MockQuerier)(nil).MapReduce), arg0, arg1)
}

// One mocks base method
func (m *MockQuerier) One(arg0 interface{}) error {
	ret := m.ctrl.Call(m, "One", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// One indicates an expected call of One
func (mr *MockQuerierMockRecorder) One(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "One", reflect.TypeOf((*MockQuerier)(nil).One), arg0)
}

// Prefetch mocks base method
func (m *MockQuerier) Prefetch(arg0 float64) Querier {
	ret := m.ctrl.Call(m, "Prefetch", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Prefetch indicates an expected call of Prefetch
func (mr *MockQuerierMockRecorder) Prefetch(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Prefetch", reflect.TypeOf((*MockQuerier)(nil).Prefetch), arg0)
}

// Select mocks base method
func (m *MockQuerier) Select(arg0 interface{}) Querier {
	ret := m.ctrl.Call(m, "Select", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Select indicates an expected call of Select
func (mr *MockQuerierMockRecorder) Select(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Select", reflect.TypeOf((*MockQuerier)(nil).Select), arg0)
}

// SetMaxScan mocks base method
func (m *MockQuerier) SetMaxScan(arg0 int) Querier {
	ret := m.ctrl.Call(m, "SetMaxScan", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// SetMaxScan indicates an expected call of SetMaxScan
func (mr *MockQuerierMockRecorder) SetMaxScan(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxScan", reflect.TypeOf((*MockQuerier)(nil).SetMaxScan), arg0)
}

// SetMaxTime mocks base method
func (m *MockQuerier) SetMaxTime(arg0 time.Duration) Querier {
	ret := m.ctrl.Call(m, "SetMaxTime", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// SetMaxTime indicates an expected call of SetMaxTime
func (mr *MockQuerierMockRecorder) SetMaxTime(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetMaxTime", reflect.TypeOf((*MockQuerier)(nil).SetMaxTime), arg0)
}

// Skip mocks base method
func (m *MockQuerier) Skip(arg0 int) Querier {
	ret := m.ctrl.Call(m, "Skip", arg0)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Skip indicates an expected call of Skip
func (mr *MockQuerierMockRecorder) Skip(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Skip", reflect.TypeOf((*MockQuerier)(nil).Skip), arg0)
}

// Snapshot mocks base method
func (m *MockQuerier) Snapshot() Querier {
	ret := m.ctrl.Call(m, "Snapshot")
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Snapshot indicates an expected call of Snapshot
func (mr *MockQuerierMockRecorder) Snapshot() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Snapshot", reflect.TypeOf((*MockQuerier)(nil).Snapshot))
}

// Sort mocks base method
func (m *MockQuerier) Sort(arg0 ...string) Querier {
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Sort", varargs...)
	ret0, _ := ret[0].(Querier)
	return ret0
}

// Sort indicates an expected call of Sort
func (mr *MockQuerierMockRecorder) Sort(arg0 ...interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sort", reflect.TypeOf((*MockQuerier)(nil).Sort), arg0...)
}

// Tail mocks base method
func (m *MockQuerier) Tail(arg0 time.Duration) *mgo_v2.Iter {
	ret := m.ctrl.Call(m, "Tail", arg0)
	ret0, _ := ret[0].(*mgo_v2.Iter)
	return ret0
}

// Tail indicates an expected call of Tail
func (mr *MockQuerierMockRecorder) Tail(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tail", reflect.TypeOf((*MockQuerier)(nil).Tail), arg0)
}
