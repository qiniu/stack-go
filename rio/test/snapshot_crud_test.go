package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotCRUD(t *testing.T) {
	testSnapShotName := "test snapname"
	createResp, err := testStack.Snapshot().SnapshotCreate(&storage.SnapshotCreateArgs{
		ZoneID:       testZoneID,
		DiskID:       &testDiskID,
		SnapshotName: testSnapShotName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, createResp)
	assert.NotEmpty(t, createResp.Data.SnapshotID)
	testSnapshotID := "8e1f7a9a-aa7a-4991-ad53-cb4845e92112"
	deleteResp, err := testStack.Snapshot().SnapshotDelete(&storage.SnapshotDeleteArgs{
		ZoneID:     testZoneID,
		SnapshotID: testSnapshotID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)
	detailResp, err := testStack.Snapshot().SnapshotDetail(&storage.SnapshotDetailArgs{
		ZoneID:     testZoneID,
		SnapshotID: testSnapshotID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, detailResp)
	assert.Equal(t, detailResp.Data.SnapshotID, testSnapshotID)
}
