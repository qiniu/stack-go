package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerDetail(t *testing.T) {
	testServerID := "42736b6d-d286-459e-b145-d113caec7382"
	detailResp, err := testStack.Server().ServerDetail(&compute.ServerDetailArgs{
		ZoneID:   testZoneID,
		ServerID: testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, detailResp)
}
