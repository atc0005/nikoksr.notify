// Code generated by mockery v2.14.0. DO NOT EDIT.

package lark

import mock "github.com/stretchr/testify/mock"

// mockSendToer is an autogenerated mock type for the sendToer type
type mockSendToer struct {
	mock.Mock
}

// SendTo provides a mock function with given fields: subject, message, id, idType
func (_m *mockSendToer) SendTo(subject string, message string, id string, idType string) error {
	ret := _m.Called(subject, message, id, idType)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string, string) error); ok {
		r0 = rf(subject, message, id, idType)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTnewMockSendToer interface {
	mock.TestingT
	Cleanup(func())
}

// newMockSendToer creates a new instance of mockSendToer. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func newMockSendToer(t mockConstructorTestingTnewMockSendToer) *mockSendToer {
	mock := &mockSendToer{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}