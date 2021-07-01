package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestEIPDetach(t *testing.T) {
	testEIPID := "531beddc-9ca4-4352-af55-dfd34e16083f"
	testServerID := "8219428a-f3b4-4959-8f42-aaf382a3ae3d"
	detachResp, err := testStack.EIP().EIPDetach(&network.EIPDetachArgs{
		ZoneID:   testZoneID,
		EIPID:    testEIPID,
		ServerID: testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, detachResp)
}
