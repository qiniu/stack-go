package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestVSwitchModify(t *testing.T) {
	testVPCID := "19ae140d-30da-4d27-82d3-65ff66868c1b"
	testVSwitchID := "8b3d39bf-8c6f-4d93-9751-ff0b01be149d"
	vswitchModifyResp, err := testStack.VSwitch().VSwitchModify(&network.VSwitchModifyArgs{
		ZoneID:      testZoneID,
		VPCID:       testVPCID,
		VSwitchID:   testVSwitchID,
		VSwitchName: "qwer",
	})
	assert.Nil(t, err)
	assert.NotNil(t, vswitchModifyResp)
}
