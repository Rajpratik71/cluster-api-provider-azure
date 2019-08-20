/*
Copyright 2019 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute/computeapi (interfaces: ResourceSkusClientAPI)

// Package mock_availabilityzones is a generated GoMock package.
package mock_availabilityzones

import (
	context "context"
	compute "github.com/Azure/azure-sdk-for-go/services/compute/mgmt/2019-07-01/compute"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockResourceSkusClientAPI is a mock of ResourceSkusClientAPI interface
type MockResourceSkusClientAPI struct {
	ctrl     *gomock.Controller
	recorder *MockResourceSkusClientAPIMockRecorder
}

// MockResourceSkusClientAPIMockRecorder is the mock recorder for MockResourceSkusClientAPI
type MockResourceSkusClientAPIMockRecorder struct {
	mock *MockResourceSkusClientAPI
}

// NewMockResourceSkusClientAPI creates a new mock instance
func NewMockResourceSkusClientAPI(ctrl *gomock.Controller) *MockResourceSkusClientAPI {
	mock := &MockResourceSkusClientAPI{ctrl: ctrl}
	mock.recorder = &MockResourceSkusClientAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockResourceSkusClientAPI) EXPECT() *MockResourceSkusClientAPIMockRecorder {
	return m.recorder
}

// List mocks base method
func (m *MockResourceSkusClientAPI) List(arg0 context.Context) (compute.ResourceSkusResultPage, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0)
	ret0, _ := ret[0].(compute.ResourceSkusResultPage)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List
func (mr *MockResourceSkusClientAPIMockRecorder) List(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockResourceSkusClientAPI)(nil).List), arg0)
}