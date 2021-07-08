package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestKeyPairDelete(t *testing.T) {
	testKeyPairName := "aaaaa"
	deleteResp, err := testStack.KeyPair().KeyPairDelete(&compute.KeyPairDeleteArgs{
		ZoneID:      testZoneID,
		KeyPairName: testKeyPairName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)
}
