package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestVPCModify(t *testing.T) {
	testVPCName := "eee"
	testVPCID := "19ae140d-30da-4d27-82d3-65ff66868c1b"
	modifyResp, err := testStack.VPC().VPCModify(&network.VPCModifyArgs{
		ZoneID:  testZoneID,
		VPCName: &testVPCName,
		VPCID:   testVPCID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, modifyResp)
}
