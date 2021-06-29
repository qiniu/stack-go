package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//ResizeDiskType 扩容磁盘类型
type ResizeDiskType string

var (
	//ResizeDiskTypeOffline ..
	ResizeDiskTypeOffline ResizeDiskType = "offline"
	//ResizeDiskTypeOnline ..
	ResizeDiskTypeOnline ResizeDiskType = "online"
)

//ResizeDiskArgs 扩容磁盘参数
type ResizeDiskArgs struct {
	ZoneID      string  `json:"zone_id"`
	DiskID      string  `json:"disk_id"`
	NewSize     int64   `json:"new_size"`
	ClientToken *string `json:"client_token,omitempty"`
}

//ResizeDiskResp 扩容磁盘返回
type ResizeDiskResp struct {
	common.Response
}

//ResizeDisk 扩容磁盘
func (d *Disk) ResizeDisk(args *ResizeDiskArgs) (resp *ResizeDiskResp, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/api/rio/v1/storage/disk/%s/resize", args.DiskID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
