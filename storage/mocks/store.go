// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/onflow/flow-emulator/storage (interfaces: Store)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/onflow/flow-emulator/types"
	state "github.com/onflow/flow-go/fvm/state"
	flow "github.com/onflow/flow-go/model/flow"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// BlockByHeight mocks base method.
func (m *MockStore) BlockByHeight(arg0 context.Context, arg1 uint64) (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByHeight", arg0, arg1)
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByHeight indicates an expected call of BlockByHeight.
func (mr *MockStoreMockRecorder) BlockByHeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByHeight", reflect.TypeOf((*MockStore)(nil).BlockByHeight), arg0, arg1)
}

// BlockByID mocks base method.
func (m *MockStore) BlockByID(arg0 context.Context, arg1 flow.Identifier) (*flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlockByID", arg0, arg1)
	ret0, _ := ret[0].(*flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlockByID indicates an expected call of BlockByID.
func (mr *MockStoreMockRecorder) BlockByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlockByID", reflect.TypeOf((*MockStore)(nil).BlockByID), arg0, arg1)
}

// CollectionByID mocks base method.
func (m *MockStore) CollectionByID(arg0 context.Context, arg1 flow.Identifier) (flow.LightCollection, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CollectionByID", arg0, arg1)
	ret0, _ := ret[0].(flow.LightCollection)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CollectionByID indicates an expected call of CollectionByID.
func (mr *MockStoreMockRecorder) CollectionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CollectionByID", reflect.TypeOf((*MockStore)(nil).CollectionByID), arg0, arg1)
}

// CommitBlock mocks base method.
func (m *MockStore) CommitBlock(arg0 context.Context, arg1 flow.Block, arg2 []*flow.LightCollection, arg3 map[flow.Identifier]*flow.TransactionBody, arg4 map[flow.Identifier]*types.StorableTransactionResult, arg5 *state.ExecutionSnapshot, arg6 []flow.Event) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CommitBlock", arg0, arg1, arg2, arg3, arg4, arg5, arg6)
	ret0, _ := ret[0].(error)
	return ret0
}

// CommitBlock indicates an expected call of CommitBlock.
func (mr *MockStoreMockRecorder) CommitBlock(arg0, arg1, arg2, arg3, arg4, arg5, arg6 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CommitBlock", reflect.TypeOf((*MockStore)(nil).CommitBlock), arg0, arg1, arg2, arg3, arg4, arg5, arg6)
}

// EventsByHeight mocks base method.
func (m *MockStore) EventsByHeight(arg0 context.Context, arg1 uint64, arg2 string) ([]flow.Event, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EventsByHeight", arg0, arg1, arg2)
	ret0, _ := ret[0].([]flow.Event)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EventsByHeight indicates an expected call of EventsByHeight.
func (mr *MockStoreMockRecorder) EventsByHeight(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EventsByHeight", reflect.TypeOf((*MockStore)(nil).EventsByHeight), arg0, arg1, arg2)
}

// LatestBlock mocks base method.
func (m *MockStore) LatestBlock(arg0 context.Context) (flow.Block, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBlock", arg0)
	ret0, _ := ret[0].(flow.Block)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBlock indicates an expected call of LatestBlock.
func (mr *MockStoreMockRecorder) LatestBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBlock", reflect.TypeOf((*MockStore)(nil).LatestBlock), arg0)
}

// LatestBlockHeight mocks base method.
func (m *MockStore) LatestBlockHeight(arg0 context.Context) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LatestBlockHeight", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// LatestBlockHeight indicates an expected call of LatestBlockHeight.
func (mr *MockStoreMockRecorder) LatestBlockHeight(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LatestBlockHeight", reflect.TypeOf((*MockStore)(nil).LatestBlockHeight), arg0)
}

// LedgerByHeight mocks base method.
func (m *MockStore) LedgerByHeight(arg0 context.Context, arg1 uint64) state.StorageSnapshot {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LedgerByHeight", arg0, arg1)
	ret0, _ := ret[0].(state.StorageSnapshot)
	return ret0
}

// LedgerByHeight indicates an expected call of LedgerByHeight.
func (mr *MockStoreMockRecorder) LedgerByHeight(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LedgerByHeight", reflect.TypeOf((*MockStore)(nil).LedgerByHeight), arg0, arg1)
}

// StoreBlock mocks base method.
func (m *MockStore) StoreBlock(arg0 context.Context, arg1 *flow.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StoreBlock", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StoreBlock indicates an expected call of StoreBlock.
func (mr *MockStoreMockRecorder) StoreBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StoreBlock", reflect.TypeOf((*MockStore)(nil).StoreBlock), arg0, arg1)
}

// TransactionByID mocks base method.
func (m *MockStore) TransactionByID(arg0 context.Context, arg1 flow.Identifier) (flow.TransactionBody, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionByID", arg0, arg1)
	ret0, _ := ret[0].(flow.TransactionBody)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionByID indicates an expected call of TransactionByID.
func (mr *MockStoreMockRecorder) TransactionByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionByID", reflect.TypeOf((*MockStore)(nil).TransactionByID), arg0, arg1)
}

// TransactionResultByID mocks base method.
func (m *MockStore) TransactionResultByID(arg0 context.Context, arg1 flow.Identifier) (types.StorableTransactionResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransactionResultByID", arg0, arg1)
	ret0, _ := ret[0].(types.StorableTransactionResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// TransactionResultByID indicates an expected call of TransactionResultByID.
func (mr *MockStoreMockRecorder) TransactionResultByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransactionResultByID", reflect.TypeOf((*MockStore)(nil).TransactionResultByID), arg0, arg1)
}
