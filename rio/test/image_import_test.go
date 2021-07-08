package test

import (
	"testing"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestImageImport(t *testing.T) {
	testImageID := "13fe0763-b78b-4f52-a0e8-c0ec16864cf4"
	testSourceURL := ""
	importResp, err := testStack.Image().ImageImport(&compute.ImageImportArgs{
		ZoneID:            testZoneID,
		ImgaeID:           testImageID,
		SourceURL:         testSourceURL,
		DestinationZoneID: "nova",
	})
	assert.Nil(t, err)
	assert.NotNil(t, importResp)
}
