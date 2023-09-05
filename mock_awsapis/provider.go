// Code generated by MockGen. DO NOT EDIT.
// Source: awsapis/provider.go

// Package mock_awsapis is a generated GoMock package.
package mock_awsapis

import (
	reflect "reflect"

	awsapis "github.com/mcastellin/aws-fail-az/awsapis"
	gomock "go.uber.org/mock/gomock"
)

// MockAWSProvider is a mock of AWSProvider interface.
type MockAWSProvider struct {
	ctrl     *gomock.Controller
	recorder *MockAWSProviderMockRecorder
}

// MockAWSProviderMockRecorder is the mock recorder for MockAWSProvider.
type MockAWSProviderMockRecorder struct {
	mock *MockAWSProvider
}

// NewMockAWSProvider creates a new mock instance.
func NewMockAWSProvider(ctrl *gomock.Controller) *MockAWSProvider {
	mock := &MockAWSProvider{ctrl: ctrl}
	mock.recorder = &MockAWSProviderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAWSProvider) EXPECT() *MockAWSProviderMockRecorder {
	return m.recorder
}

// NewAutoScalingApi mocks base method.
func (m *MockAWSProvider) NewAutoScalingApi() awsapis.AutoScalingApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewAutoScalingApi")
	ret0, _ := ret[0].(awsapis.AutoScalingApi)
	return ret0
}

// NewAutoScalingApi indicates an expected call of NewAutoScalingApi.
func (mr *MockAWSProviderMockRecorder) NewAutoScalingApi() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewAutoScalingApi", reflect.TypeOf((*MockAWSProvider)(nil).NewAutoScalingApi))
}

// NewDynamodbApi mocks base method.
func (m *MockAWSProvider) NewDynamodbApi() awsapis.DynamodbApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewDynamodbApi")
	ret0, _ := ret[0].(awsapis.DynamodbApi)
	return ret0
}

// NewDynamodbApi indicates an expected call of NewDynamodbApi.
func (mr *MockAWSProviderMockRecorder) NewDynamodbApi() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewDynamodbApi", reflect.TypeOf((*MockAWSProvider)(nil).NewDynamodbApi))
}

// NewEc2Api mocks base method.
func (m *MockAWSProvider) NewEc2Api() awsapis.Ec2Api {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEc2Api")
	ret0, _ := ret[0].(awsapis.Ec2Api)
	return ret0
}

// NewEc2Api indicates an expected call of NewEc2Api.
func (mr *MockAWSProviderMockRecorder) NewEc2Api() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEc2Api", reflect.TypeOf((*MockAWSProvider)(nil).NewEc2Api))
}

// NewEcsApi mocks base method.
func (m *MockAWSProvider) NewEcsApi() awsapis.EcsApi {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewEcsApi")
	ret0, _ := ret[0].(awsapis.EcsApi)
	return ret0
}

// NewEcsApi indicates an expected call of NewEcsApi.
func (mr *MockAWSProviderMockRecorder) NewEcsApi() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewEcsApi", reflect.TypeOf((*MockAWSProvider)(nil).NewEcsApi))
}
