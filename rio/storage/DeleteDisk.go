package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// DeleteDiskArgs 删除磁盘参数
type DeleteDiskArgs struct {
	ZoneID string `json:"zone_id"`
	DiskID string `json:"disk_id"`
}

// DeleteDiskResp 删除磁盘返回
type DeleteDiskResp struct {
	common.Response
}

// DeleteDisk 删除磁盘
func (d *Disk) DeleteDisk(args *DeleteDiskArgs) (resp *DeleteDiskResp, err error) {
	url := fmt.Sprintf("%s/disk/%s", StorageURLPrefix, args.DiskID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
