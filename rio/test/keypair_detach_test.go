package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestKeyPairDetach(t *testing.T) {
	testKeyPairName := "aaa"
	testServerID := "5ed06d35-fe9b-4947-87f3-87769c61d305"
	detachResp, err := testStack.KeyPair().KeyPairDetach(&compute.KeyPairDetachArgs{
		ZoneID:      testZoneID,
		KeyPairName: testKeyPairName,
		ServerID:    testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, detachResp)
}
