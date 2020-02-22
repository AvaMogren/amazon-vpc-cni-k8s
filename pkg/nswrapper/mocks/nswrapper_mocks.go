// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-vpc-cni-k8s/pkg/nswrapper (interfaces: NS)

// Package mock_nswrapper is a generated GoMock package.
package mock_nswrapper

import (
	reflect "reflect"

	ns "github.com/containernetworking/cni/pkg/ns"
	gomock "github.com/golang/mock/gomock"
)

// MockNS is a mock of NS interface
type MockNS struct {
	ctrl     *gomock.Controller
	recorder *MockNSMockRecorder
}

// MockNSMockRecorder is the mock recorder for MockNS
type MockNSMockRecorder struct {
	mock *MockNS
}

// NewMockNS creates a new mock instance
func NewMockNS(ctrl *gomock.Controller) *MockNS {
	mock := &MockNS{ctrl: ctrl}
	mock.recorder = &MockNSMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockNS) EXPECT() *MockNSMockRecorder {
	return m.recorder
}

// WithNetNSPath mocks base method
func (m *MockNS) WithNetNSPath(arg0 string, arg1 func(ns.NetNS) error) error {
	ret := m.ctrl.Call(m, "WithNetNSPath", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithNetNSPath indicates an expected call of WithNetNSPath
func (mr *MockNSMockRecorder) WithNetNSPath(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithNetNSPath", reflect.TypeOf((*MockNS)(nil).WithNetNSPath), arg0, arg1)
}
