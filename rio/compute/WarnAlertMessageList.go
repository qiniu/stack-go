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
	ZoneID         string                  `json:"header:x-rio-zone-id"`
	InstanceIP     string                  `json:"query:instance_ip"`
	InstanceName   string                  `json:"query:instance_name"`
	LevelList      []common.WarnAlertLevel `json:"query:level_list"`
	ResourceTypeID *string                 `json:"query:resource_type_id"`
	PageNum        *int64                  `json:"query:page_num"`
	PageSize       *int64                  `json:"query:page_size"`
}

// WarnDescribesAlertMessageResp 查询告警信息返回
type WarnDescribesAlertMessageResp struct {
	common.Response
	Results []*AlertMessageDetail `json:"results"`
}

// WarnDescribesAlertMessage 告警修改
func (d *Warn) WarnDescribesAlertMessage(args *WarnDescribesAlertMessageArgs) (resp *WarnDescribesAlertMessageResp, err error) {
	str := "/api/rio/v1/compute/warn"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/alert")).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
