package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//SnapshotModifyArgs 修改快照参数
type SnapshotModifyArgs struct {
	ZoneID     string `json:"zone_id"`
	SnapshotID string `json:"snapshot_id:"`
	// 快照的显示名json度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	// 为防止和自动快照的名称冲突，不能以auto开头。
	SnapshotName string `json:"snapshot_name,omitempty"`
	// 快照的描述。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。 默认值：空
	Description string `json:"description,omitempty"`
}

//SnapshotModifyResp 修改快照返回
type SnapshotModifyResp struct {
	common.Response
}

//SnapshotModify 修改快照
func (d *Snapshot) SnapshotModify(args *SnapshotModifyArgs) (resp *SnapshotModifyResp, err error) {
	req := client.NewRequest(http.MethodPatch, fmt.Sprintf("/api/rio/v1/storage/snapshot/%s", args.SnapshotID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	//	req := client.NewRequest(http.MethodPatch, fmt.Sprintf("/api/rio/v1/storage/snapshot/%s", args.SnapshotID)).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
