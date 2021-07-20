package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnDeleteArgs 删除参数
type WarnDeleteArgs struct {
	ZoneID string `json:"zone_id"`
	WarnID string `json:"id"`
}

// WarnDeleteResp 删除返回
type WarnDeleteResp struct {
	common.Response
}

// WarnDelete 告警删除
func (w *Warn) WarnDelete(args *WarnDeleteArgs) (resp *WarnDeleteResp, err error) {
	url := fmt.Sprintf("%s/warn/%s", ComputURLPrefix, args.WarnID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = w.client.Call(req, &resp)
	return
}
