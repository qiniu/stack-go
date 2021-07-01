package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/network"
	"github.com/stretchr/testify/assert"
)

func TestEIPModify(t *testing.T) {
	testEIPID := "10db31d9-db84-4f0e-b4b6-fd552b3650c3"
	testEIPName := "kkk"
	modifyResp, err := testStack.EIP().EIPModify(&network.EIPModifyArgs{
		ZoneID:    testZoneID,
		EIPID:     testEIPID,
		EIPName:   testEIPName,
		Bandwidth: 2,
	})
	assert.Nil(t, err)
	assert.NotNil(t, modifyResp)
}
