package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestSecurityGroupDelete(t *testing.T) {
	testSecurityGroupID := "feceed42-1400-4020-84a5-accac8515960"
	deleteResp, err := testStack.SecurityGroup().SecurityGroupDelete(&compute.SecurityGroupDeleteArgs{
		ZoneID:          testZoneID,
		SecurityGroupID: testSecurityGroupID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)
}
