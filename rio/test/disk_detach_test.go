package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func Test_DiskDetach(t *testing.T) {
	testDiskID := "a403aaf4-8d18-41f8-8c58-8bffd019b6c0"
	testServerID := "6bcfcb3b-25fa-4b6e-aac4-cd02f68ba65d"
	detachResp, err := testStack.Disk().DetachDisk(&storage.DetachDiskArgs{
		ZoneID:   testZoneID,
		DiskID:   testDiskID,
		ServerID: testServerID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, detachResp)
}
