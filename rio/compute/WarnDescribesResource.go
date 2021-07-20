package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnInstanceListArgs 查询参数
type WarnInstanceListArgs struct {
	ZoneID string `json:"zone_id"`
	WarnID string `json:"warn_id"`
}

// WarnInstanceResp 查询返回
type WarnInstanceResp struct {
	common.Response
	Data []*WarnInfo `json:"data"`
}

// WarnInstance 告警实例
func (w *Warn) WarnInstance(args *WarnInstanceListArgs) (resp *WarnInstanceResp, err error) {
	url := fmt.Sprintf("%s/warn/instances", ComputURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = w.client.Call(req, &resp)
	return
}
