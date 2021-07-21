package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// CreateDiskArgs 创建磁盘参数
type CreateDiskArgs struct {
	ZoneID      string  `json:"zone_id"`
	Size        *int64  `json:"size,omitempty"`
	DiskName    *string `json:"disk_name,omitempty"`
	Description *string `json:"description,omitempty"`
	SnapshotID  *string `json:"snapshot_id,omitempty"`
}

// CreateDiskResp 创建磁盘返回
type CreateDiskResp struct {
	common.Response
	Data struct {
		DiskID string `json:"disk_id"`
	} `json:"data"`
}

// CreateDisk 创建磁盘
func (d *Disk) CreateDisk(args *CreateDiskArgs) (resp *CreateDiskResp, err error) {
	url := fmt.Sprintf("%s/disk", StorageURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
