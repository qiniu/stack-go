package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeImagesParams 镜像列表参数
type DescribeImagesParams struct {
	Page            *int             `json:"page"`
	Size            *int             `json:"size"`
	RegionID        *string          `json:"region_id"`
	ImageOwnerAlias *ImageOwnerAlias `json:"image_owner_alias"`
	ImageID         *string          `json:"image_id"`
	ImageName       *string          `json:"image_name"`
	InstanceID      *string          `json:"instance_id"`
	OSName          *string          `json:"os_name"`
}

// DescribeImagesResponse 镜像列表返回数据
type DescribeImagesResponse struct {
	Page      int     `json:"page"`
	Size      int     `json:"size"`
	Total     int     `json:"total"`
	RequestID string  `json:"request_id"`
	Data      []Image `json:"data"`
}

// DescribeImages 镜像列表
func (s *ECS) DescribeImages(p *DescribeImagesParams) (resp *DescribeImagesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/image").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
