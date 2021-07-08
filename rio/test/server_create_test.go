package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerCreate(t *testing.T) {
	testImageID := "13fe0763-b78b-4f52-a0e8-c0ec16864cf4"
	testFlavorID := "26d48369-c91e-41d1-949f-75a483fe8f09"
	testSecurityGroupID := "2fb40beb-5ba2-45c2-851a-5f914a6e96c9"
	testVSwitchID := "d5440a3d-16c8-4086-b73c-d6c49c9f725f"
	testVPCID := "098bce6d-bf2d-4705-8c1f-f3209017a590"
	testDiskName := "hhh"
	testDataDisks := []compute.ServerDiskArgs{}
	createResp, err := testStack.Server().ServerCreate(&compute.ServerCreateArgs{
		ZoneID:          testZoneID,
		ImageID:         testImageID,
		FlavorID:        testFlavorID,
		SecurityGroupID: testSecurityGroupID,
		VSwitchID:       &testVSwitchID,
		ServerName:      "sss",
		Password:        "123456",
		KeyPairName:     "dddddddddd",
		VPCID:           &testVPCID,
		SystemDisk: compute.ServerDiskArgs{
			Size:     40,
			DiskName: &testDiskName,
		},
		DataDisks: testDataDisks,
	})
	assert.Nil(t, err)
	assert.NotNil(t, createResp)
}
