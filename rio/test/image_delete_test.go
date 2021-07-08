package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestImageDelete(t *testing.T) {
	testImageID := "87a603fe-ee76-4950-9bbb-24dd03568f81"
	deleteResp, err := testStack.Image().ImageDelete(&compute.ImageDeleteArgs{
		ZoneID:  testZoneID,
		ImageID: testImageID,
	})
	assert.Nil(t, err)
	assert.NotNil(t, deleteResp)
}
