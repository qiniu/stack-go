package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//DetachDiskArgs 解绑磁盘参数
type DetachDiskArgs struct {
	ZoneID   string `json:"zone_id"`
	DiskID   string `json:"disk_id"`
	ServerID string `json:"server_id"`
}

//DetachDiskResp 解绑磁盘返回
type DetachDiskResp struct {
	common.Response
}

//DetachDisk 解绑磁盘
func (d *Disk) DetachDisk(args *DetachDiskArgs) (resp *DetachDiskResp, err error) {
	url := fmt.Sprintf("%s/disk/%s/detach", StorageURLPrefix, args.DiskID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
