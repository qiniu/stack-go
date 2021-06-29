package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotList(t *testing.T) {
	testSnapshotID := "eeec2544-1234-4469-95d0-52eb8875dc5e"
	listResp, err := testStack.Snapshot().SnapshotList(&storage.SnapshotListArgs{
		ZoneID:     testZoneID,
		SnapshotID: &testSnapshotID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, listResp)
}
