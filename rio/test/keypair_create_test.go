package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestKeyPairCreate(t *testing.T) {
	createResp, err := testStack.KeyPair().KeyPairCreate(&compute.KeyPairCreateArgs{
		ZoneID:      testZoneID,
		KeyPairName: "abb-a",
	})

	assert.Nil(t, err)
	assert.NotNil(t, createResp)
}
