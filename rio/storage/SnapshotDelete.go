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
}

//SnapshotDeleteResp 删除快照返回
type SnapshotDeleteResp struct {
	common.Response
}

//SnapshotDelete 删除快照
func (d *Snapshot) SnapshotDelete(args *SnapshotDeleteArgs) (resp *SnapshotCreateResp, err error) {
	url := fmt.Sprintf("%s/snapshot/%s", StorageURLPrefix)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
