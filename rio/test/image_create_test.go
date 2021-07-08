package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestImageCreate(t *testing.T) {
	testImageName := "qqq"
	testSnapshotID := "87a603fe-ee76-4950-9bbb-24dd03568f82"
	testServerID := "5ed06d35-fe9b-4947-87f3-87769c61d305"
	createResp, err := testStack.Image().ImageCreate(&compute.ImageCreateArgs{
		ZoneID:     testZoneID,
		SnapshotID: testSnapshotID,
		ServerID:   testServerID,
		ImageName:  &testImageName,
	})
	assert.Nil(t, err)
	assert.NotNil(t, createResp)
}
