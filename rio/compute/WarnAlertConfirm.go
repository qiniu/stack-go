package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnAlertConfirmArgs 处理参数
type WarnAlertConfirmArgs struct {
	ZoneID    string    `json:"zone_id"`
	AlertList []*string `json:"alert_list" json:"alert_list"`
	ISDelete  bool      `json:"is_delete"  json:"is_delete"`
	Confirm   *string   `json:"confirm"    json:"confirm"`
}

// WarnAlertConfirmResp 处理参数
type WarnAlertConfirmResp struct {
	common.Response
}

// WarnAlertConfirm 告警修改
func (d *Warn) WarnAlertConfirm(args *WarnAlertConfirmArgs) (resp *WarnAlertConfirmResp, err error) {
	str := "/api/rio/v1/compute/warn"
	req := client.NewRequest(http.MethodPut, fmt.Sprintf(str+"/alert")).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
