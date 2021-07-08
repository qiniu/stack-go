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
}

// WarnInstance .
func (d *Warn) WarnInstance(args *WarnInstanceListArgs) (resp *WarnInstanceResp, err error) {
	str := "/api/rio/v1/compute/warn"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/instances")).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
