package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestVSwitchCRUD(t *testing.T) {
	testVPCID := "19ae140d-30da-4d27-82d3-65ff66868c1b"
	testVSwitchName := "az"
	testVSwitchCidr := "129.168.0.0/24"
	vswitchCreateResp, err := testStack.VSwitch().VSwitchCreate(&network.VSwitchCreateArgs{
		ZoneID:      testZoneID,
		VPCID:       testVPCID,
		VSwitchName: &testVSwitchName,
		VSwitchCidr: testVSwitchCidr,
	})
	assert.Nil(t, err)
	assert.NotNil(t, vswitchCreateResp)
	testVSwitchID := "f8778280-9eb6-44f6-9fc8-a0ce424a3f4c"
	vswitchDeleteResp, err := testStack.VSwitch().VSwitchDelete(&network.VSwitchDeleteArgs{
		ZoneID:    testZoneID,
		VSwitchID: testVSwitchID,
		VPCID:     testVPCID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, vswitchDeleteResp)
}
