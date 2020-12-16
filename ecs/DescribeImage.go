package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeImageParams 镜像详情参数
type DescribeImageParams struct {
	ImageID  string `json:"image_id"`
	RegionID string `json:"region_id"`
}

// DescribeImageResponse 镜像详情返回数据
type DescribeImageResponse struct {
	RequestID string `json:"request_id"`
	Data      Image  `json:"data"`
}

// DescribeImage 镜像详情
func (s *ECS) DescribeImage(p *DescribeImageParams) (resp *DescribeImageResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/vm/image/%s", p.ImageID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
