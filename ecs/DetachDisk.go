package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DetachDiskParams 磁盘卸载参数
type DetachDiskParams struct {
	// TODO
}

// DetachDiskResponse 磁盘卸载返回数据
type DetachDiskResponse struct {
	// TODO
}

// DetachDisk 磁盘卸载
func (s *ECS) DetachDisk(p *DetachDiskParams) (resp *DetachDiskResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/disk/:id/attach")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
