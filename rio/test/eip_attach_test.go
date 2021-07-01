package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestEIPAttach(t *testing.T) {
	testEIPID := "10db31d9-db84-4f0e-b4b6-fd552b3650c3"
	testServerID := "39f19536-c894-4132-b3a5-bad942b1c1b1"
	testPrivateIPAddress := "192.168.0.5"
	attachResp, err := testStack.EIP().EIPAttach(&network.EIPAttachArgs{
		ZoneID:           testZoneID,
		EIPID:            testEIPID,
		ServerID:         testServerID,
		PrivateIPAddress: testPrivateIPAddress,
	})
	assert.Nil(t, err)
	assert.NotNil(t, attachResp)
}
