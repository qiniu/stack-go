package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/storage"
	"github.com/stretchr/testify/assert"
)

func TestSnapshotModify(t *testing.T) {
	testSnapshotID := "eeec2544-1234-4469-95d0-52eb8875dc5e"
	modifyResp, err := testStack.Snapshot().SnapshotModify(&storage.SnapshotModifyArgs{
		ZoneID:       testZoneID,
		SnapshotID:   testSnapshotID,
		SnapshotName: "aaa",
	})
	assert.Nil(t, err)
	assert.NotNil(t, modifyResp)
}
