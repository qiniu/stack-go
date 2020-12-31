package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeImagesParams 镜像列表参数
type DescribeImagesParams struct {

	// 页码。默认：1
	Page *int `json:"page"`

	// 分页大小，最大100。默认20
	Size *int `json:"size"`

	// 实例所属的地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	RegionID *string `json:"region_id"`

	// 镜像来源。取值范围：
	// - `system` ：七牛云提供的公共镜像。
	// - `self` ：您创建的自定义镜像。
	// - `others` ：其他七牛云用户共享给您的镜像。
	// - `marketplace` ：镜像市场提供的镜像。您查询到的云市场镜像可以直接使用，无需提前订阅。您需要自行留意云市场镜像的收费详情。
	// 默认值：空，空表示返回取值为system、self以及others的结果。
	ImageOwnerAlias *ImageOwnerAlias `json:"image_owner_alias"`

	// 镜像ID。
	ImageID *string `json:"image_id"`

	// 镜像名称。
	//
	// 长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 默认值：空
	ImageName *string `json:"image_name"`

	// 实例 ID
	InstanceID *string `json:"instance_id"`

	// 操作系统的中文显示名称。
	OSName *string `json:"os_name"`
}

// DescribeImagesResponse 镜像列表返回数据
type DescribeImagesResponse struct {
	Page      int     `json:"page"`       // 页码
	Size      int     `json:"size"`       // 分页大小
	Total     int     `json:"total"`      // 镜像数量
	RequestID string  `json:"request_id"` // 请求 ID
	Data      []Image `json:"data"`       // 镜像列表
}

// DescribeImages 镜像列表
//
// @GSD:URI GET /v1/vm/image
func (s *ECS) DescribeImages(p *DescribeImagesParams) (resp *DescribeImagesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/image").WithRegionID(p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
