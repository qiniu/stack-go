package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteDiskParams 磁盘删除参数
type DeleteDiskParams struct {
	RegionID string `json:"region_id"`
	DiskID   string `json:"disk_id"`
}

// DeleteDiskResponse 磁盘删除返回数据
type DeleteDiskResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteDisk 磁盘删除
func (s *ECS) DeleteDisk(p *DeleteDiskParams) (resp *DeleteDiskResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/disk/%s", p.DiskID)).WithRegionID(&p.RegionID)
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
