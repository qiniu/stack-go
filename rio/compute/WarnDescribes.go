package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnListArgs 查询参数
type WarnListArgs struct {
	ZoneID   string `json:"zone-id"`
	WarnName string `json:"warn_name"`
	PageSize int64  `json:"page_size"`
	PageNum  int64  `json:"page_num"`
}

// WarnListResp 查询返回
type WarnListResp struct {
	common.Response
	Total int64 // 数据总量
}

// WarnList 告警列表
func (d *Warn) WarnList(args *WarnListArgs) (resp *WarnListResp, err error) {
	str := "/api/rio/v1/compute/warn"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str)).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
