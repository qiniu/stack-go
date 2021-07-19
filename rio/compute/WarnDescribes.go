package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnListArgs 查询参数
type WarnListArgs struct {
	ZoneID   string `json:"zone_id"`
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
func (w *Warn) WarnList(args *WarnListArgs) (resp *WarnListResp, err error) {
	url := fmt.Sprintf("%s/warn", ComputURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = w.client.Call(req, &resp)
	return
}
