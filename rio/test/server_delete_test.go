package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerDelete(t *testing.T) {
	testServerID := "75a4f9a0-a2aa-44ae-b9fe-39210604b127"
	deleteResp, err := testStack.Server().ServerDelete(&compute.ServerDeleteArgs{
		ZoneID:   testZoneID,
		ServerID: testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)
}
