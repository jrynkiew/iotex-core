// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/types/types.go

// Package mock_apitypes is a generated GoMock package.
package mock_apitypes

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	apitypes "github.com/iotexproject/iotex-core/api/types"
	block "github.com/iotexproject/iotex-core/blockchain/block"
)

// MockWeb3ResponseWriter is a mock of Web3ResponseWriter interface.
type MockWeb3ResponseWriter struct {
	ctrl     *gomock.Controller
	recorder *MockWeb3ResponseWriterMockRecorder
}

// MockWeb3ResponseWriterMockRecorder is the mock recorder for MockWeb3ResponseWriter.
type MockWeb3ResponseWriterMockRecorder struct {
	mock *MockWeb3ResponseWriter
}

// NewMockWeb3ResponseWriter creates a new mock instance.
func NewMockWeb3ResponseWriter(ctrl *gomock.Controller) *MockWeb3ResponseWriter {
	mock := &MockWeb3ResponseWriter{ctrl: ctrl}
	mock.recorder = &MockWeb3ResponseWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWeb3ResponseWriter) EXPECT() *MockWeb3ResponseWriterMockRecorder {
	return m.recorder
}

// Write mocks base method.
func (m *MockWeb3ResponseWriter) Write(arg0 interface{}) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write.
func (mr *MockWeb3ResponseWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockWeb3ResponseWriter)(nil).Write), arg0)
}

// MockResponder is a mock of Responder interface.
type MockResponder struct {
	ctrl     *gomock.Controller
	recorder *MockResponderMockRecorder
}

// MockResponderMockRecorder is the mock recorder for MockResponder.
type MockResponderMockRecorder struct {
	mock *MockResponder
}

// NewMockResponder creates a new mock instance.
func NewMockResponder(ctrl *gomock.Controller) *MockResponder {
	mock := &MockResponder{ctrl: ctrl}
	mock.recorder = &MockResponderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockResponder) EXPECT() *MockResponderMockRecorder {
	return m.recorder
}

// Exit mocks base method.
func (m *MockResponder) Exit() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Exit")
}

// Exit indicates an expected call of Exit.
func (mr *MockResponderMockRecorder) Exit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exit", reflect.TypeOf((*MockResponder)(nil).Exit))
}

// Respond mocks base method.
func (m *MockResponder) Respond(arg0 string, arg1 *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Respond", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Respond indicates an expected call of Respond.
func (mr *MockResponderMockRecorder) Respond(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Respond", reflect.TypeOf((*MockResponder)(nil).Respond), arg0, arg1)
}

// MockListener is a mock of Listener interface.
type MockListener struct {
	ctrl     *gomock.Controller
	recorder *MockListenerMockRecorder
}

// MockListenerMockRecorder is the mock recorder for MockListener.
type MockListenerMockRecorder struct {
	mock *MockListener
}

// NewMockListener creates a new mock instance.
func NewMockListener(ctrl *gomock.Controller) *MockListener {
	mock := &MockListener{ctrl: ctrl}
	mock.recorder = &MockListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockListener) EXPECT() *MockListenerMockRecorder {
	return m.recorder
}

// AddResponder mocks base method.
func (m *MockListener) AddResponder(arg0 apitypes.Responder) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddResponder", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddResponder indicates an expected call of AddResponder.
func (mr *MockListenerMockRecorder) AddResponder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddResponder", reflect.TypeOf((*MockListener)(nil).AddResponder), arg0)
}

// ReceiveBlock mocks base method.
func (m *MockListener) ReceiveBlock(arg0 *block.Block) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReceiveBlock", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// ReceiveBlock indicates an expected call of ReceiveBlock.
func (mr *MockListenerMockRecorder) ReceiveBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReceiveBlock", reflect.TypeOf((*MockListener)(nil).ReceiveBlock), arg0)
}

// RemoveResponder mocks base method.
func (m *MockListener) RemoveResponder(arg0 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveResponder", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RemoveResponder indicates an expected call of RemoveResponder.
func (mr *MockListenerMockRecorder) RemoveResponder(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveResponder", reflect.TypeOf((*MockListener)(nil).RemoveResponder), arg0)
}

// Start mocks base method.
func (m *MockListener) Start() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start.
func (mr *MockListenerMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockListener)(nil).Start))
}

// Stop mocks base method.
func (m *MockListener) Stop() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop.
func (mr *MockListenerMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockListener)(nil).Stop))
}