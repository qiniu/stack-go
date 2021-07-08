package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// MonitorDiskDataArgs disk监控参数
type MonitorDiskDataArgs struct {
	ZoneID    string `json:"zone_id"`
	DiskID    string `json:"disk_id"`
	ServerID  string `json:"server_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// MonitorDiskDataResp 监控数据result
type MonitorDiskDataResp struct {
	common.Response
}

// MonitorDiskData ..
func (d *Monitor) MonitorDiskData(args *MonitorDiskDataArgs) (resp *MonitorDiskDataResp, err error) {
	str := "/api/rio/v1/compute/monitor"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/disk/%s", args.DiskID)).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
