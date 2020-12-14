package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeImagesParams 镜像列表参数
type DescribeImagesParams struct {
	// TODO
}

// DescribeImagesResponse 镜像列表返回数据
type DescribeImagesResponse struct {
	// TODO
}

// DescribeImages 镜像列表
func (s *ECS) DescribeImages(p *DescribeImagesParams) (resp *DescribeImagesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/image")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
