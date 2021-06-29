package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

//DetachDiskArgs 解绑磁盘参数
type DetachDiskArgs struct {
	ZoneID   string `json:"zone_id"`
	DiskID   string `json:"disk_id"`
	ServerID string `json:"server_id"`
}

//DetachDiskResp 解绑磁盘返回
type DetachDiskResp struct {
	// common.Response
}

//DetachDisk 解绑磁盘
func (d *Disk) DetachDisk(args *DetachDiskArgs) (resp *DetachDiskResp, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/api/rio/v1/storage/disk/%s/detach", args.DiskID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
