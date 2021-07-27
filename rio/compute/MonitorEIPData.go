package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// MonitorEipDataArgs eip监控参数
type MonitorEipDataArgs struct {
	ZoneID    string `json:"zone_id"`
	EipID     string `json:"eip_id"`
	ServerID  string `json:"server_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// MonitorEipDataResp eip监控数据返回
type MonitorEipDataResp struct {
	common.Response
	Data []*EipMonitorDataInfo `json:"data"`
}

// MonitorEipData eip监控
func (m *Monitor) MonitorEipData(args *MonitorEipDataArgs) (resp *MonitorEipDataResp, err error) {
	url := fmt.Sprintf("%s/monitor/eip/%s", ComputURLPrefix, args.EipID)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = m.client.Call(req, &resp)
	return
}
