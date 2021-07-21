package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SnapshotModifyArgs 修改快照参数
type SnapshotModifyArgs struct {
	ZoneID       string `json:"zone_id"`
	SnapshotID   string `json:"snapshot_id:"`
	SnapshotName string `json:"snapshot_name,omitempty"`
	Description  string `json:"description,omitempty"`
}

// SnapshotModifyResp 修改快照返回
type SnapshotModifyResp struct {
	common.Response
}

// SnapshotModify 修改快照
func (s *Snapshot) SnapshotModify(args *SnapshotModifyArgs) (resp *SnapshotModifyResp, err error) {
	url := fmt.Sprintf("%s/snapshot", StorageURLPrefix)
	req := client.NewRequest(http.MethodPatch, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
