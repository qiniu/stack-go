package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// AttachDiskParams 磁盘挂载参数
type AttachDiskParams struct {
	// TODO
}

// AttachDiskResponse 磁盘挂载返回数据
type AttachDiskResponse struct {
	// TODO
}

// AttachDisk 磁盘挂载
func (s *ECS) AttachDisk(p *AttachDiskParams) (resp *AttachDiskResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/disk/:id/attach")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
