package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DetachDiskParams 磁盘卸载参数
type DetachDiskParams struct {
	RegionID   string `json:"region_id"`
	DiskID     string `json:"disk_id"`
	InstanceID string `json:"instance_id"`
}

// DetachDiskResponse 磁盘卸载返回数据
type DetachDiskResponse struct {
	RequestID string `json:"request_id"`
}

// DetachDisk 磁盘卸载
func (s *ECS) DetachDisk(p *DetachDiskParams) (resp *DetachDiskResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/disk/%s/attach", p.DiskID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
