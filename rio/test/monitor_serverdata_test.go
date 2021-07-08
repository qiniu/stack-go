package test

import (
	"testing"
	"time"

	"github.com/qiniu/stack-go/rio/compute"
	"github.com/stretchr/testify/assert"
)

func TestMonitorServerData(t *testing.T) {
	testServerID := "5ed06d35-fe9b-4947-87f3-87769c61d305"
	testStartTime := time.Now().AddDate(0, 0, -1).Unix() * 1000
	testEndTime := time.Now().Unix() * 1000
	monitorserverResp, err := testStack.Monitor().MonitorServerData(&compute.MonitorServerDataArgs{
		ZoneID:    testZoneID,
		ServerID:  testServerID,
		StartTime: testStartTime,
		EndTime:   testEndTime,
	})
	assert.Nil(t, err)
	assert.NotNil(t, monitorserverResp)
}
