package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestDiskReset(t *testing.T) {
	testDiskID := "f54fe222-840d-465e-a66a-2080ab12f2db"
	testSnapshotID := "eeec2544-1234-4469-95d0-52eb8875dc5e"
	resetReso, err := testStack.Disk().ResetDisk(&storage.ResetDiskArgs{
		ZoneID:     testZoneID,
		DiskID:     testDiskID,
		SnapshotID: testSnapshotID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, resetReso)
	//assert.Equal(t, resetReso.DiskID, testDiskID)

}
