package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateDiskParams 磁盘创建参数
type CreateDiskParams struct {
	// TODO
}

// CreateDiskResponse 磁盘创建返回数据
type CreateDiskResponse struct {
	// TODO
}

// CreateDisk 磁盘创建
func (s *ECS) CreateDisk(p *CreateDiskParams) (resp *CreateDiskResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/disk")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
