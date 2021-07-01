package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestEIPList(t *testing.T) {
	listResp, err := testStack.EIP().EIPList(&network.EIPListArgs{
		ZoneID: testZoneID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, listResp)
}
