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

	DiskID       *string `json:"disk_id"`
	SnapshotName string  `json:"snapshot_name,omitempty"`
	Description  *string `json:"description,omitempty"`
	ClientToken  *string `json:"client_token,omitempty"`
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
	url := fmt.Sprintf("%s/snapshot", StorageURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
