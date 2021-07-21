package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnDescribesAlertMessageArgs 查询告警信息参数
type WarnDescribesAlertMessageArgs struct {
	// 可用区 ID
	ZoneID         string                  `json:"zone_id"`
	InstanceIP     string                  `json:"instance_ip"`
	InstanceName   string                  `json:"instance_name"`
	LevelList      []common.WarnAlertLevel `json:"level_list"`
	ResourceTypeID *string                 `json:"resource_type_id"`
	PageNum        *int64                  `json:"page_num"`
	PageSize       *int64                  `json:"page_size"`
}

// WarnDescribesAlertMessageResp 查询告警信息返回
type WarnDescribesAlertMessageResp struct {
	common.Response
	Results []*AlertMessageDetail `json:"results"`
}

// WarnDescribesAlertMessage 告警修改
func (w *Warn) WarnDescribesAlertMessage(args *WarnDescribesAlertMessageArgs) (resp *WarnDescribesAlertMessageResp, err error) {
	url := fmt.Sprintf("%s/warn/alert", ComputURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = w.client.Call(req, &resp)
	return
}
