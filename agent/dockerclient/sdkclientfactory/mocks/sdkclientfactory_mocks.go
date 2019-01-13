// Copyright 2015-2019 Amazon.com, Inc. or its affiliates. All Rights Reserved.
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
// Source: github.com/aws/amazon-ecs-agent/agent/dockerclient/sdkclientfactory (interfaces: Factory)

// Package mock_sdkclientfactory is a generated GoMock package.
package mock_sdkclientfactory

import (
	reflect "reflect"

	dockerclient "github.com/aws/amazon-ecs-agent/agent/dockerclient"
	sdkclient "github.com/aws/amazon-ecs-agent/agent/dockerclient/sdkclient"
	gomock "github.com/golang/mock/gomock"
)

// MockFactory is a mock of Factory interface
type MockFactory struct {
	ctrl     *gomock.Controller
	recorder *MockFactoryMockRecorder
}

// MockFactoryMockRecorder is the mock recorder for MockFactory
type MockFactoryMockRecorder struct {
	mock *MockFactory
}

// NewMockFactory creates a new mock instance
func NewMockFactory(ctrl *gomock.Controller) *MockFactory {
	mock := &MockFactory{ctrl: ctrl}
	mock.recorder = &MockFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockFactory) EXPECT() *MockFactoryMockRecorder {
	return m.recorder
}

// FindClientAPIVersion mocks base method
func (m *MockFactory) FindClientAPIVersion(arg0 sdkclient.Client) dockerclient.DockerVersion {
	ret := m.ctrl.Call(m, "FindClientAPIVersion", arg0)
	ret0, _ := ret[0].(dockerclient.DockerVersion)
	return ret0
}

// FindClientAPIVersion indicates an expected call of FindClientAPIVersion
func (mr *MockFactoryMockRecorder) FindClientAPIVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindClientAPIVersion", reflect.TypeOf((*MockFactory)(nil).FindClientAPIVersion), arg0)
}

// FindKnownAPIVersions mocks base method
func (m *MockFactory) FindKnownAPIVersions() []dockerclient.DockerVersion {
	ret := m.ctrl.Call(m, "FindKnownAPIVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// FindKnownAPIVersions indicates an expected call of FindKnownAPIVersions
func (mr *MockFactoryMockRecorder) FindKnownAPIVersions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindKnownAPIVersions", reflect.TypeOf((*MockFactory)(nil).FindKnownAPIVersions))
}

// FindSupportedAPIVersions mocks base method
func (m *MockFactory) FindSupportedAPIVersions() []dockerclient.DockerVersion {
	ret := m.ctrl.Call(m, "FindSupportedAPIVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// FindSupportedAPIVersions indicates an expected call of FindSupportedAPIVersions
func (mr *MockFactoryMockRecorder) FindSupportedAPIVersions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindSupportedAPIVersions", reflect.TypeOf((*MockFactory)(nil).FindSupportedAPIVersions))
}

// GetClient mocks base method
func (m *MockFactory) GetClient(arg0 dockerclient.DockerVersion) (sdkclient.Client, error) {
	ret := m.ctrl.Call(m, "GetClient", arg0)
	ret0, _ := ret[0].(sdkclient.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetClient indicates an expected call of GetClient
func (mr *MockFactoryMockRecorder) GetClient(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClient", reflect.TypeOf((*MockFactory)(nil).GetClient), arg0)
}

// GetDefaultClient mocks base method
func (m *MockFactory) GetDefaultClient() (sdkclient.Client, error) {
	ret := m.ctrl.Call(m, "GetDefaultClient")
	ret0, _ := ret[0].(sdkclient.Client)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDefaultClient indicates an expected call of GetDefaultClient
func (mr *MockFactoryMockRecorder) GetDefaultClient() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDefaultClient", reflect.TypeOf((*MockFactory)(nil).GetDefaultClient))
}
