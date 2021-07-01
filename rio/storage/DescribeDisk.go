package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// DescribeDiskArgs 查询磁盘参数
type DescribeDiskArgs struct {
	ZoneID string `json:"zone_id"`
	DiskID string `json:"disk_id"`
}

// DescribeDiskResp 查询磁盘返回
type DescribeDiskResp struct {
	common.Response
	Data *DiskInfo `json:"data"`
}

// DescribeDisk 查询磁盘详情
func (d *Disk) DescribeDisk(args *DescribeDiskArgs) (resp *DescribeDiskResp, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/api/rio/v1/storage/disk/%s", args.DiskID)).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
