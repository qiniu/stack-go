package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestSecurityGroupCreate(t *testing.T) {
	testSecurityGroupName := "aaa"
	createResp, err := testStack.SecurityGroup().SecurityGroupCreate(&compute.SecurityGroupCreateArgs{
		ZoneID:            testZoneID,
		SecurityGroupName: &testSecurityGroupName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, createResp)
}
