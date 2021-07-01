package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestEIPCRUD(t *testing.T) {
	testEIPName := "test openapi"
	createResp, err := testStack.EIP().EIPCreate(&network.EIPCreateArgs{
		ZoneID:    testZoneID,
		EIPName:   &testEIPName,
		Bandwidth: 1,
	})
	assert.Nil(t, err)
	assert.NotNil(t, createResp)
	testEIPID := "f2ede89d-62c4-4d41-bebc-ac5942cfdbed"
	deleteResp, err := testStack.EIP().EIPDelete(&network.EIPDeleteArgs{
		ZoneID: testZoneID,
		EIPID:  testEIPID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)
}
