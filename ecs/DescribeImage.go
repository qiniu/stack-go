package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeImageParams 镜像详情参数
type DescribeImageParams struct {
	// TODO
}

// DescribeImageResponse 镜像详情返回数据
type DescribeImageResponse struct {
	// TODO
}

// DescribeImage 镜像详情
func (s *ECS) DescribeImage(p *DescribeImageParams) (resp *DescribeImageResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/image/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
