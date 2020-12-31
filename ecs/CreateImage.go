package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// CreateImageParams 镜像创建参数
type CreateImageParams struct {
	// 镜像所在的地域 ID。 您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	RegionID string `json:"region_id"`

	// 根据指定的快照创建自定义镜像
	//
	SnapshotID *string `json:"snapshot_id"`

	// 实例 ID
	//
	// InstanceID 与 SnapshotID 必须传入一个
	InstanceID *string `json:"instance_id"`

	// 镜像名称。
	//
	// 长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 默认值：空
	ImageName *string `json:"image_name"`

	// 镜像版本号
	ImageVersion *string `json:"image_version"`

	// 镜像的描述信息。长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。
	// 默认值：空
	Description *string `json:"description"`

	// 保证请求幂等性， 客户端生成的一个唯一值
	ClientToken *string `json:"client_token"`

	// 计费参数（必须）
	*params.CostParams
}

// CreateImageResponse 镜像创建返回数据
type CreateImageResponse struct {
	RequestID string `json:"request_id"`
	Data      Image  `json:"data"`
}

// CreateImage 镜像创建
func (s *ECS) CreateImage(p *CreateImageParams) (resp *CreateImageResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/image").WithJSONBody(p).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
