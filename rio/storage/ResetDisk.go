package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//ResetDiskArgs 重置磁盘参数
type ResetDiskArgs struct {
	ZoneID string `json:"zone_id"`
	DiskID string `json:"disk_id"`
	// 需要恢复到某一磁盘阶段的历史快照ID。
	SnapshotID string `json:"snapshot_id"`
}

//ResetDiskResp 重置磁盘返回
type ResetDiskResp struct {
	common.Response
}

//ResetDisk 重置磁盘
func (d *Disk) ResetDisk(args *ResetDiskArgs) (resp *ResetDiskResp, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/api/rio/v1/storage/disk/%s/reset", args.DiskID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
