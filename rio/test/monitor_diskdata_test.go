package test

import (
	"fmt"
	"testing"
	"time"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestMonitorDiskData(t *testing.T) {
	testDiskID := "7833c256-6516-4a68-893a-e731c7714216"
	testServerID := "5ed06d35-fe9b-4947-87f3-87769c61d305"
	testStarTime := time.Now().AddDate(0, 0, -1).Unix() * 1000
	testEndTime := time.Now().Unix() * 1000
	monitordiskResp, err := testStack.Monitor().MonitorDiskData(&compute.MonitorDiskDataArgs{
		ZoneID:   testZoneID,
		DiskID:   testDiskID,
		ServerID: testServerID,

		StartTime: testStarTime,
		EndTime:   testEndTime,
	})

	fmt.Println(testStarTime, testEndTime)
	assert.Nil(t, err)
	assert.NotNil(t, monitordiskResp)
}
