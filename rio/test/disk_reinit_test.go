package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestDiskReinit(t *testing.T) {
	testDiskID := "a403aaf4-8d18-41f8-8c58-8bffd019b6c0"
	reinitResp, err := testStack.Disk().ReInitDisk(&storage.ReInitDiskArgs{
		ZoneID: testZoneID,
		DiskID: testDiskID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, reinitResp)
}
