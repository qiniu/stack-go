package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//ModifyDiskArgs 修改磁盘参数
type ModifyDiskArgs struct {
	ZoneID           string  `json:"zone_id"`
	DiskID           string  `json:"disk_id"`
	DiskName         *string `json:"disk_name,omitempty"`
	Description      *string `json:"description,omitempty"`
	DeleteWithServer *bool   `json:"delete_with_server,omitempty"`
}

//ModifyDiskResp 修改磁盘返回
type ModifyDiskResp struct {
	common.Response
}

//ModifyDisk 修改磁盘
func (d *Disk) ModifyDisk(args *ModifyDiskArgs) (resp *ModifyDiskResp, err error) {
	url := fmt.Sprintf("%s/disk/%s", StorageURLPrefix, args.DiskID)
	req := client.NewRequest(http.MethodPatch, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
