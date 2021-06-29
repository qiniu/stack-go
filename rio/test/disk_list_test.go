package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestDiskList(t *testing.T) {
	listResp, err := testStack.Disk().ListDisk(&storage.ListDiskArgs{
		ZoneID: testZoneID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, listResp)
}
