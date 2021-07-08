package test

import (
	"testing"
	"time"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestMonitorEIPData(t *testing.T) {
	testStartTime := time.Now().AddDate(0, 0, -1).Unix() * 1000
	testEndTime := time.Now().Unix() * 1000
	testEipID := "0c60ac76-f130-4899-9558-f0e42c39414d"
	testServerID := "lb-0145b1b1-190a-4a7a-b2f7-e4b043661fae"
	monitoreipResp, err := testStack.Monitor().MonitorEipData(&compute.MonitorEipDataArgs{
		ZoneID:    testZoneID,
		EipID:     testEipID,
		ServerID:  testServerID,
		StartTime: testStartTime,
		EndTime:   testEndTime,
	})
	assert.Nil(t, err)
	assert.NotNil(t, monitoreipResp)
}
