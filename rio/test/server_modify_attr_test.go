package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerModifyAttr(t *testing.T) {
	testServerName := "asda"
	attrResp, err := testStack.Server().ServerModifyAttr(&compute.ServerModifyAttrArgs{
		ZoneID:     testZoneID,
		ServerID:   "42736b6d-d286-459e-b145-d113caec7382",
		ServerName: testServerName,
		Password:   "123456",
	})
	assert.Nil(t, err)
	assert.NotNil(t, attrResp)
}
