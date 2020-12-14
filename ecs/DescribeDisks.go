package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeDisksParams 磁盘列表参数
type DescribeDisksParams struct {
	Page       *int      `json:"page"`
	Size       *int      `json:"size"`
	DiskID     *string   `json:"disk_id"`
	DiskName   *string   `json:"disk_name"`
	InstanceID *string   `json:"instance_id"`
	ZoneID     *string   `json:"zone_id"`
	Type       *DiskType `json:"type"`
	Tag        *string   `json:"tag"`
	RegionID   *string   `json:"region_id"`
}

// DescribeDisksResponse 磁盘列表返回数据
type DescribeDisksResponse struct {
	Page      int    `json:"page"`
	Size      int    `json:"size"`
	Total     int    `json:"total"`
	RequestID string `json:"request_id"`
	Data      []Disk `json:"data"`
}

// DescribeDisks 磁盘列表
func (s *ECS) DescribeDisks(p *DescribeDisksParams) (resp *DescribeDisksResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/disk").WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
