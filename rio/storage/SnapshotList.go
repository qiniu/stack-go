package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SnapshotListArgs 快照列表参数
type SnapshotListArgs struct {
	ZoneID         string           `json:"zone_id"`
	SnapshotID     *string          `json:"snapshot_id"`
	SnapshotName   *string          `json:"snapshot_name"`
	DiskID         *string          `json:"disk_id"`
	ServerID       *string          `json:"server_id"`
	SourceDiskType *common.DiskType `json:"source_disk_type"`
}

// SnapshotListResp 快照列表返回
type SnapshotListResp struct {
	common.Response
	Data []*SnapshotInfo `json:"data"`
}

// SnapshotList 快照列表
func (s *Snapshot) SnapshotList(args *SnapshotListArgs) (resp *SnapshotListResp, err error) {
	url := fmt.Sprintf("%s/snapshot", StorageURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
