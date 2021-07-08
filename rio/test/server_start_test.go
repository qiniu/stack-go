package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestServerStart(t *testing.T) {
	testServerID := "967d385f-8621-4bdd-accc-bf155860b28f"
	startResp, err := testStack.Server().ServerStart(&compute.ServerStartArgs{
		ZoneID:   testZoneID,
		ServerID: testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, startResp)
}
