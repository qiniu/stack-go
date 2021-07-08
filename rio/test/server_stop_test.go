package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerStop(t *testing.T) {
	testServerID := "967d385f-8621-4bdd-accc-bf155860b28f"
	stopResp, err := testStack.Server().ServerStop(&compute.ServerStopArgs{
		ZoneID:   testZoneID,
		ServerID: testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, stopResp)
}
