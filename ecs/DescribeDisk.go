package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeDiskParams 磁盘详情参数
type DescribeDiskParams struct {
	RegionID string `json:"region_id"` // 镜像所在的地域 ID。 您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	DiskID   string `json:"disk_id"`   // 磁盘 ID
}

// DescribeDiskResponse 磁盘详情返回数据
type DescribeDiskResponse struct {
	RequestID string `json:"request_id"`
	Data      Disk   `json:"data"`
}

// DescribeDisk 磁盘详情
func (s *ECS) DescribeDisk(p *DescribeDiskParams) (resp *DescribeDiskResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/disk/%s", p.DiskID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
