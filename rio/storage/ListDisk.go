package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//ListDiskArgs 磁盘列表参数
type ListDiskArgs struct {
	ZoneID   string                 `json:"zone_id"`
	DiskID   *string                `json:"disk_id"`
	DiskName *string                `json:"disk_name"`
	ServerID *string                `json:"server_id"`
	DiskType *common.DiskType       `json:"disk_type"`
	Status   *common.DiskStatusType `json:"status"`
}

//ListDiskResp 磁盘列表返回
type ListDiskResp struct {
	common.Response
	Data []*DiskInfo `json:"data"`
}

//ListDisk 磁盘列表
func (d *Disk) ListDisk(args *ListDiskArgs) (resp *ListDiskResp, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/api/rio/v1/storage/disk")).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
