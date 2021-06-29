package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestDiskResize(t *testing.T) {

	testDiskID := "a403aaf4-8d18-41f8-8c58-8bffd019b6c0"
	resizeResp, err := testStack.Disk().ResizeDisk(&storage.ResizeDiskArgs{
		ZoneID:  testZoneID,
		DiskID:  testDiskID,
		NewSize: 52,
	})
	assert.Nil(t, err)
	assert.NotNil(t, resizeResp)
}
