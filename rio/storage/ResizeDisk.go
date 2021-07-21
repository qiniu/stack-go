package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ResizeDiskType 扩容磁盘类型
type ResizeDiskType string

var (
	// ResizeDiskTypeOffline 扩容
	ResizeDiskTypeOffline ResizeDiskType = "offline"
	// ResizeDiskTypeOnline 扩容
	ResizeDiskTypeOnline ResizeDiskType = "online"
)

// ResizeDiskArgs 扩容磁盘参数
type ResizeDiskArgs struct {
	ZoneID      string  `json:"zone_id"`
	DiskID      string  `json:"disk_id"`
	NewSize     int64   `json:"new_size"`
	ClientToken *string `json:"client_token,omitempty"`
}

// ResizeDiskResp 扩容磁盘返回
type ResizeDiskResp struct {
	common.Response
}

// ResizeDisk 扩容磁盘
func (d *Disk) ResizeDisk(args *ResizeDiskArgs) (resp *ResizeDiskResp, err error) {
	url := fmt.Sprintf("%s/disk/%s/resize", StorageURLPrefix, args.DiskID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
