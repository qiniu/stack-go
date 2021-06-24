package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func Test_DiskCRUD(t *testing.T) {
	// Create
	testDiskSize := int64(20)
	testDiskName := "test openapi"
	createResp, err := testStack.Disk().CreateDisk(&storage.CreateDiskArgs{
		ZoneID:   testZoneID,
		Size:     &testDiskSize,
		DiskName: &testDiskName,
	})

	assert.Nil(t, err)
	assert.NotNil(t, createResp)
	assert.NotEmpty(t, createResp.Data.DiskID)

	testDiskID := createResp.Data.DiskID

	// Update TODO

	// Retrive
	describeResp, err := testStack.Disk().DescribeDisk(&storage.DescribeDiskArgs{
		ZoneID: testZoneID,
		DiskID: testDiskID,
	})

	assert.Nil(t, err)
	assert.NotNil(t, describeResp)
	assert.Equal(t, describeResp.Data.DiskID, testDiskID)

	// Delete
	deleteResp, err := testStack.Disk().DeleteDisk(&storage.DeleteDiskArgs{
		ZoneID: testZoneID,
		DiskID: testDiskID,
	})

	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)

	// Verify Delete
	describeResp, err = testStack.Disk().DescribeDisk(&storage.DescribeDiskArgs{
		ZoneID: testZoneID,
		DiskID: testDiskID,
	})

	assert.Nil(t, err)
	assert.NotNil(t, describeResp)
}
