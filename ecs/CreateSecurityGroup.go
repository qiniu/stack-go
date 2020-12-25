package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateSecurityGroupParams 安全组创建参数
type CreateSecurityGroupParams struct {
	// 安全组所属地域ID。您可以调用DescribeRegions查看最新的七牛云地域列表。
	RegionID string `json:"region_id"`

	// 安全组名称。
	//
	// 长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http://和https://开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。默认值：空。
	SecurityGroupName *string `json:"security_group_name"`

	// 安全组描述信息。长度为2~256个英文或中文字符，不能以http://和https://开头。默认值：空。
	Description *string `json:"description"`

	// 安全组所属VPC ID
	VPCID *string `json:"vpc_id"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`
}

// CreateSecurityGroupResponse 安全组创建返回数据
type CreateSecurityGroupResponse struct {
	RequestID string `json:"request_id"`
}

// CreateSecurityGroup 创建安全组
func (s *ECS) CreateSecurityGroup(p *CreateSecurityGroupParams) (resp *CreateSecurityGroupResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/sgr").WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
