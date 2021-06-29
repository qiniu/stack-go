package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//SnapshotCreateArgs 创建快照参数
type SnapshotCreateArgs struct {
	ZoneID string `json:"zone_id"`

	// 云盘ID
	DiskID *string `json:"disk_id"`
	// 快照的显示名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 为防止和自动快照的名称冲突，不能以auto开头。
	SnapshotName string `json:"snapshot_name,omitempty"`
	// 快照的描述。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。 默认值：空
	Description *string `json:"description,omitempty"`
	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token,omitempty"`
}

//SnapshotCreateResp 创建快照返回
type SnapshotCreateResp struct {
	common.Response
	Data struct {
		SnapshotID *string `json:"snapshot_id"`
	} `json:"data"`
}

//SnapshotCreate 创建快照
func (d *Snapshot) SnapshotCreate(args *SnapshotCreateArgs) (resp *SnapshotCreateResp, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/api/rio/v1/storage/snapshot")).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
