package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerLeaveSecurity(t *testing.T) {
	testServerID := "c2f4895c-3435-4281-a461-85c47325fb11"
	testSecurityGroupID := "2fb40beb-5ba2-45c2-851a-5f914a6e96c9"
	leaveResp, err := testStack.Server().ServerLeaveSecurityGroup(&compute.ServerLeaveSecurityGroupArgs{
		ZoneID:          testZoneID,
		ServerID:        testServerID,
		SecurityGroupID: testSecurityGroupID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, leaveResp)
}
