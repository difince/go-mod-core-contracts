//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package responses

import (
	"testing"

	"github.com/edgexfoundry/go-mod-core-contracts/v2/dtos"
	"github.com/stretchr/testify/assert"
)

func TestNewDeviceResponse(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedMessage := "unit test message"
	expectedDevice := dtos.Device{Name: "test device"}
	actual := NewDeviceResponse(expectedRequestId, expectedMessage, expectedStatusCode, expectedDevice)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedMessage, actual.Message)
	assert.Equal(t, expectedDevice, actual.Device)
}

func TestNewDeviceResponseNoMessage(t *testing.T) {
	expectedRequestId := "123456"
	expectedStatusCode := 200
	expectedDevice := dtos.Device{Name: "test device"}
	actual := NewDeviceResponseNoMessage(expectedRequestId, expectedStatusCode, expectedDevice)

	assert.Equal(t, expectedRequestId, actual.RequestId)
	assert.Equal(t, expectedStatusCode, actual.StatusCode)
	assert.Equal(t, expectedDevice, actual.Device)
}
