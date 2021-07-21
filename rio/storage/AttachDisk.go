package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// AttachDiskArgs 绑定磁盘参数
type AttachDiskArgs struct {
	ZoneID           string `json:"zone_id"`
	DiskID           string `json:"disk_id"`
	ServerID         string `json:"server_id,omitempty"`
	DeleteWithServer *bool  `json:"delete_with_server,omitempty"`
}

//AttachDiskResp 绑定磁盘返回
type AttachDiskResp struct {
	common.Response
}

// AttachDisk 绑定磁盘
func (d *Disk) AttachDisk(args *AttachDiskArgs) (resp *AttachDiskResp, err error) {
	url := fmt.Sprintf("%s/disk/%s/attach", StorageURLPrefix, args.DiskID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}