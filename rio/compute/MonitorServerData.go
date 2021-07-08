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

// MonitorServerDataResp 返回
type MonitorServerDataResp struct {
	common.Response
}

// MonitorServerData 主机监控
func (d *Monitor) MonitorServerData(args *MonitorServerDataArgs) (resp *MonitorServerDataResp, err error) {
	str := "/api/rio/v1/compute/monitor"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/server/%s", args.ServerID)).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}