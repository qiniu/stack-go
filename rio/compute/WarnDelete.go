package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnDeleteArgs 删除参数
type WarnDeleteArgs struct {
	ZoneID string `json:"header:x-rio-zone-id"`
	WarnID string `json:"path:id"`
}

// WarnDeleteResp 删除返回
type WarnDeleteResp struct {
	common.Response
}

// WarnDelete 告警删除
func (d *Warn) WarnDelete(args *WarnDeleteArgs) (resp *WarnDeleteResp, err error) {
	str := "/api/rio/v1/compute/warn"
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf(str+"%s", args.WarnID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
