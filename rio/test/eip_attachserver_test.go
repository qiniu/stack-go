package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestEIPAttachServer(t *testing.T) {
	testEIPID := "10db31d9-db84-4f0e-b4b6-fd552b3650c3"
	attachserverResp, err := testStack.EIP().FindEIPAttachServer(&network.FindEIPAttachServerArgs{
		ZoneID: testZoneID,
		EIPID:  testEIPID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, attachserverResp)
}
