package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//ReInitDiskArgs 初始化磁盘参数
type ReInitDiskArgs struct {
	ZoneID string `json:"zone_id"`
	DiskID string `pos:"disk_id"`
}

//ReInitDiskResp 初始化磁盘返回
type ReInitDiskResp struct {
	common.Response
}

//ReInitDisk 初始化磁盘
func (d *Disk) ReInitDisk(args *ReInitDiskArgs) (resp *ReInitDiskResp, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/api/rio/v1/storage/disk/%s/reinit", args.DiskID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
