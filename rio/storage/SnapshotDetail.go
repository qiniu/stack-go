package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//SnapshotDetailArgs 查询快照参数
type SnapshotDetailArgs struct {
	ZoneID     string `json:"zone_id"`
	SnapshotID string `json:"snapshot_id"`
}

//SnapshotDetailResp 查询快照返回
type SnapshotDetailResp struct {
	common.Response
	Data *SnapshotInfo `json:"data"`
}

//SnapshotDetail 查询快照详情
func (d *Snapshot) SnapshotDetail(args *SnapshotDetailArgs) (resp *SnapshotDetailResp, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/api/rio/v1/storage/snapshot/%s", args.SnapshotID)).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return

}
