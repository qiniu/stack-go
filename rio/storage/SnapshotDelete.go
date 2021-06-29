package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//SnapshotDeleteArgs 删除快照参数
type SnapshotDeleteArgs struct {
	ZoneID     string `json:"zone_id"`
	SnapshotID string `json:"snapshot_id"`
	// 是否强制删除有磁盘关联的快照。说明 删除后该磁盘无法重新初始化。
	Force *bool `json:"force,omitempty"` // TODO 暂不支持
}

//SnapshotDeleteResp 删除快照返回
type SnapshotDeleteResp struct {
	common.Response
}

//SnapshotDelete 删除快照
func (d *Snapshot) SnapshotDelete(args *SnapshotDeleteArgs) (resp *SnapshotCreateResp, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/api/rio/v1/storage/snapshot/%s", args.SnapshotID)).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
