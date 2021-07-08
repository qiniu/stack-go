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

// MonitorEipDataResp 监控数据result
type MonitorEipDataResp struct {
	common.Response
}

// MonitorEipData 监控数据
func (d *Monitor) MonitorEipData(args *MonitorEipDataArgs) (resp *MonitorEipDataResp, err error) {
	str := "/api/rio/v1/compute/monitor"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/eip/%s", args.EipID)).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
