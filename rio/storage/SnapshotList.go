package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//SnapshotListArgs 快照列表参数
type SnapshotListArgs struct {
	// TODO 暂不支持分页
	// Marker string `pos:"query:marker"`
	// Limit  int    `pos:"query:limit"`

	ZoneID         string           `json:"zone_id"`
	SnapshotID     *string          `json:"snapshot_id"`
	SnapshotName   *string          `json:"snapshot_name"`
	DiskID         *string          `json:"disk_id"`
	ServerID       *string          `json:"server_id"`
	SourceDiskType *common.DiskType `json:"source_disk_type"`
}

//SnapshotListResp 快照列表返回
type SnapshotListResp struct {
	common.Response
	Data []*SnapshotInfo `json:"data"`
}

//SnapshotList 快照列表
func (d *Snapshot) SnapshotList(args *SnapshotListArgs) (resp *SnapshotListResp, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/api/rio/v1/storage/snapshot")).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
