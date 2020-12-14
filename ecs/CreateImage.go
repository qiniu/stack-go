package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateImageParams 镜像创建参数
type CreateImageParams struct {
	// TODO
}

// CreateImageResponse 镜像创建返回数据
type CreateImageResponse struct {
	// TODO
}

// CreateImage 镜像创建
func (s *ECS) CreateImage(p *CreateImageParams) (resp *CreateImageResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/image")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
