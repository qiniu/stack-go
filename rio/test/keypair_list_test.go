package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestKeyPairList(t *testing.T) {
	TestKeyPairName := "zzz"
	listResp, err := testStack.KeyPair().KeyPairList(&compute.KeyPairListArgs{
		ZoneID:      testZoneID,
		KeyPairName: TestKeyPairName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, listResp)
}
