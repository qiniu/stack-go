package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestKeyPairDetail(t *testing.T) {
	testKeyPairName := "qqq"
	detailResp, err := testStack.KeyPair().KeyPairDetail(&compute.KeyPairDetailArgs{
		ZoneID:      testZoneID,
		KeyPairName: testKeyPairName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, detailResp)
}
