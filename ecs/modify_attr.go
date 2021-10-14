package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyInstanceAttribute .
type ModifyInstanceAttribute struct {
	Action     string `ason:"Action"     json:"-"`           // 系统规定参数。取值：ModifyInstanceAttribute
	InstanceId string `ason:"InstanceId" json:"instance_id"` // 实例ID。

	// 实例所属的地域ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	RegionId string `ason:"RegionId" json:"region_id"`

	// 实例名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	InstanceName *string `ason:"InstanceName,omitempty" json:"instance_name"`
}

// ModifyInstanceAttributeResponse .
type ModifyInstanceAttributeResponse struct {
	RequestID string `json:"request_id"`
}

// ModifyInstanceAttribute .
func (s *ECS) ModifyInstanceAttribute(p *ModifyInstanceAttribute) (resp *AttachDiskResponse, err error) {
	req := client.NewRequest(http.MethodPut, fmt.Sprintf("/v1/vm/instance/%s/attribute", p.InstanceId)).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
