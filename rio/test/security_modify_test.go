package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestSecurityGroupModify(t *testing.T) {
	testSecurityGroupID := "2fb40beb-5ba2-45c2-851a-5f914a6e96c9"
	testSecurityGroupName := "aaa"
	modifyResp, err := testStack.SecurityGroup().SecurityGroupModify(&compute.SecurityGroupModifyArgs{
		ZoneID:            testZoneID,
		SecurityGroupID:   testSecurityGroupID,
		SecurityGroupName: testSecurityGroupName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, modifyResp)
}
