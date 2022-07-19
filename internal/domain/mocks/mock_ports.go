// Code generated by MockGen. DO NOT EDIT.
// Source: internal/domain/ports.go

// Package mock_domain is a generated GoMock package.
package mock_domain

import (
	reflect "reflect"
	domain "train-reservation/internal/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockBookingReferenceProvider is a mock of BookingReferenceProvider interface.
type MockBookingReferenceProvider struct {
	ctrl     *gomock.Controller
	recorder *MockBookingReferenceProviderMockRecorder
}

// MockBookingReferenceProviderMockRecorder is the mock recorder for MockBookingReferenceProvider.
type MockBookingReferenceProviderMockRecorder struct {
	mock *MockBookingReferenceProvider
}

// NewMockBookingReferenceProvider creates a new mock instance.
func NewMockBookingReferenceProvider(ctrl *gomock.Controller) *MockBookingReferenceProvider {
	mock := &MockBookingReferenceProvider{ctrl: ctrl}
	mock.recorder = &MockBookingReferenceProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBookingReferenceProvider) EXPECT() *MockBookingReferenceProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockBookingReferenceProvider) Get() (domain.BookingReference, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get")
	ret0, _ := ret[0].(domain.BookingReference)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockBookingReferenceProviderMockRecorder) Get() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockBookingReferenceProvider)(nil).Get))
}

// MockTrainProvider is a mock of TrainProvider interface.
type MockTrainProvider struct {
	ctrl     *gomock.Controller
	recorder *MockTrainProviderMockRecorder
}

// MockTrainProviderMockRecorder is the mock recorder for MockTrainProvider.
type MockTrainProviderMockRecorder struct {
	mock *MockTrainProvider
}

// NewMockTrainProvider creates a new mock instance.
func NewMockTrainProvider(ctrl *gomock.Controller) *MockTrainProvider {
	mock := &MockTrainProvider{ctrl: ctrl}
	mock.recorder = &MockTrainProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTrainProvider) EXPECT() *MockTrainProviderMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockTrainProvider) Get(trainId domain.TrainId) (domain.Train, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", trainId)
	ret0, _ := ret[0].(domain.Train)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get.
func (mr *MockTrainProviderMockRecorder) Get(trainId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockTrainProvider)(nil).Get), trainId)
}

// MockReservationClient is a mock of ReservationClient interface.
type MockReservationClient struct {
	ctrl     *gomock.Controller
	recorder *MockReservationClientMockRecorder
}

// MockReservationClientMockRecorder is the mock recorder for MockReservationClient.
type MockReservationClientMockRecorder struct {
	mock *MockReservationClient
}

// NewMockReservationClient creates a new mock instance.
func NewMockReservationClient(ctrl *gomock.Controller) *MockReservationClient {
	mock := &MockReservationClient{ctrl: ctrl}
	mock.recorder = &MockReservationClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReservationClient) EXPECT() *MockReservationClientMockRecorder {
	return m.recorder
}

// Reserve mocks base method.
func (m *MockReservationClient) Reserve(train domain.Train, confirmation domain.ReservationConfirmation) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Reserve", train, confirmation)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Reserve indicates an expected call of Reserve.
func (mr *MockReservationClientMockRecorder) Reserve(train, confirmation interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Reserve", reflect.TypeOf((*MockReservationClient)(nil).Reserve), train, confirmation)
}
