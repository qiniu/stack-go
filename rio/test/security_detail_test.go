package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestSecurityGroupDetail(t *testing.T) {
	testSecurityGroupID := "2fb40beb-5ba2-45c2-851a-5f914a6e96c9"
	detailResp, err := testStack.SecurityGroup().SecurityGroupDetail(&compute.SecurityGroupDetailArgs{
		ZoneID:          testZoneID,
		SecurityGroupID: testSecurityGroupID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, detailResp)
}
