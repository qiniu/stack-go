package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerVNC(t *testing.T) {
	testServerID := "c2f4895c-3435-4281-a461-85c47325fb11"
	vncResp, err := testStack.Server().ServerVNC(&compute.ServerVNCArgs{
		ZoneID:   testZoneID,
		ServerID: testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, vncResp)
}
