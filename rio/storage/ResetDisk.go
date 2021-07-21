package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ResetDiskArgs 重置磁盘参数
type ResetDiskArgs struct {
	ZoneID     string `json:"zone_id"`
	DiskID     string `json:"disk_id"`
	SnapshotID string `json:"snapshot_id"`
}

// ResetDiskResp 重置磁盘返回
type ResetDiskResp struct {
	common.Response
}

// ResetDisk 重置磁盘
func (d *Disk) ResetDisk(args *ResetDiskArgs) (resp *ResetDiskResp, err error) {
	url := fmt.Sprintf("%s/disk/%s/reset", StorageURLPrefix, args.DiskID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
