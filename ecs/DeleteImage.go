package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteImageParams 镜像删除参数
type DeleteImageParams struct {

	// 自定义镜像所在的地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	RegionID string `json:"region_id"`

	// 镜像ID。如果指定的自定义镜像不存在，则请求将被忽略。
	ImageID string `json:"image_id"`

	// 是否执行强制删除。取值范围：
	// - `true`：强制删除自定义镜像，忽略当前镜像是否被其他实例使用。
	// - `false`：正常删除自定义镜像，删除前检查当前镜像是否被其他实例使用。
	// 默认值：`false`
	Force *bool `json:"force"`

	// 是否删除相关快照，默认不删除
	DeleteRelatedSnapshot bool `json:"delete_related_snapshot"`
}

// DeleteImageResponse 镜像删除返回数据
type DeleteImageResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteImage 镜像删除
func (s *ECS) DeleteImage(p *DeleteImageParams) (resp *DeleteImageResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/image/%s", p.ImageID)).WithJSONBody(p).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
