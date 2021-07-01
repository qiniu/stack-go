package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestVPCCRUD(t *testing.T) {
	testCIDR := "192.169.0.0/16"
	testVPCName := "qko"
	testVSwitchName := "S-S307G"
	testVSwitchCidr := "192.168.0.0/24"
	//testVSwitchName:="S-abcA"
	createResp, err := testStack.VPC().VPCCreate(&network.CreateVPCWithVSwitchArgs{
		ZoneID: testZoneID,
		VPC: &network.VPCCreateArgs{
			VPCName: &testVPCName,
			CIDR:    &testCIDR,
		},
		VSwitch: &network.VSwitchCreateArgs{
			VSwitchName: &testVSwitchName,
			VSwitchCidr: testVSwitchCidr,
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, createResp)
	testVPCID := "26b718f3-a08c-4d8c-8664-a488a54a8c47"
	deleteResp, err := testStack.VPC().VPCDelete(&network.VPCDeleteArgs{
		ZoneID: testZoneID,
		VPCID:  testVPCID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)
}
