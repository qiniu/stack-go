package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestDiskModify(t *testing.T) {
	testDiskID := "e693f34e-4d24-4555-9bbf-9646836293d3"
	modifyResp, err := testStack.Disk().ModifyDisk(&storage.ModifyDiskArgs{
		ZoneID: testZoneID,
		DiskID: testDiskID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, modifyResp)
}
