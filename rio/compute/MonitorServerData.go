package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// MonitorServerDataArgs 主机监控参数
type MonitorServerDataArgs struct {
	ZoneID    string `json:"zone_id"`
	ServerID  string `json:"server_id"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

// MonitorServerDataResp 主机监控返回
type MonitorServerDataResp struct {
	common.Response
}

// MonitorServerData 主机监控
func (m *Monitor) MonitorServerData(args *MonitorServerDataArgs) (resp *MonitorServerDataResp, err error) {
	url := fmt.Sprintf("%s/monitor/server/%s", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = m.client.Call(req, &resp)
	return
}
