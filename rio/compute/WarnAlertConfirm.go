package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnAlertConfirmArgs 告警确认参数
type WarnAlertConfirmArgs struct {
	ZoneID    string    `json:"zone_id"`
	AlertList []*string `json:"alert_list"`
	ISDelete  bool      `json:"is_delete"`
	Confirm   *string   `json:"confirm"`
}

// WarnAlertConfirmResp 告警确认返回
type WarnAlertConfirmResp struct {
	common.Response
}

// WarnAlertConfirm 告警确认
func (w *Warn) WarnAlertConfirm(args *WarnAlertConfirmArgs) (resp *WarnAlertConfirmResp, err error) {
	url := fmt.Sprintf("%s/warn/alert", ComputURLPrefix)
	req := client.NewRequest(http.MethodPut, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = w.client.Call(req, &resp)
	return
}
