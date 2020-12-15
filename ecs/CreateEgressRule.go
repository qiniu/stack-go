package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateEgressRuleParams 安全组出网规则创建参数
type CreateEgressRuleParams struct {
	SecurityGroupID string   `json:"security_group_id"`
	RegionID        string   `json:"region_id"`
	AuthType        AuthType `json:"auth_type"`
	AuthList        []string `json:"auth_list"`       // 授权列表
	PortRangeList   []string `json:"port_range_list"` // 端口范围列表

	// 传输层协议。取值大小写敏感。取值范围：
	// - icmp
	// - gre
	// - tcp
	// - udp
	// - all：支持所有协议
	IPProtocol IPProtocol `json:"ip_protocol"` //enum tcp | udp | icmp | gre | all

	// 目的端安全组开放的传输层协议相关的端口范围。取值范围：
	// - TCP/UDP协议：取值范围为1~65535。使用斜线（/）隔开起始端口和终止端口。例如：1/200
	// - ICMP协议：-1/-1
	// - GRE协议：-1/-1
	// - all：-1/-1
	PortRange string `json:"port_range"`

	// 需要设置访问权限的目的端安全组ID。
	// - 至少设置一项DestGroupId或者DestCidrIp参数。
	// - 如果指定了DestGroupId没有指定参数DestCidrIp，则参数NicType取值只能为intranet。
	// - 如果同时指定了DestGroupId和DestCidrIp，则默认以DestCidrIp为准。
	DestGroupID *string `json:"dest_group_id"`

	// 跨账户设置安全组规则时，目的端安全组所属的七牛云账户。
	// - 如果DestGroupOwnerAccount及DestGroupOwnerId均未设置，则认为是设置您其他安全组的访问权限。
	// - 如果已经设置参数DestCidrIp，则参数DestGroupOwnerAccount无效。
	DestGroupOwnerAccount *string `json:"dest_group_owner_account"`

	// 跨账户设置安全组规则时，目的端安全组所属的七牛云账户ID。
	// - 如果DestGroupOwnerId及DestGroupOwnerAccount均未设置，则认为是设置您其他安全组的访问权限。
	// - 如果您已经设置参数DestCidrIp，则参数DestGroupOwnerId无效。
	DestGroupOwnerID *string `json:"dest_group_owner_id"`

	// 目的端IP地址范围。支持CIDR格式和IPv4格式的IP地址范围。默认值：无
	DestCidrIP *string `json:"dest_cidr_ip"` //IPv4 only

	// 	源端IP地址范围。支持CIDR格式和IPv4格式的IP地址范围。默认值：无
	SourceCidrIP *string `json:"source_cidr_ip"` // IPv4 only, default 0.0.0.0/0

	// 源端安全组开放的传输层协议相关的端口范围。取值范围：
	// - TCP/UDP协议：1~65535。使用斜线（/）隔开起始端口和终止端口。例如：1/200
	// - ICMP协议：-1/-1
	// - GRE协议：-1/-1
	// - all：-1/-1
	SourcePortRange *string `json:"source_port_range"`

	// 设置访问权限。取值范围：
	// - accept：接受访问
	// - drop：拒绝访问，不返回拒绝信息
	//
	// 默认值：accept
	Policy *PermissionPolicy `json:"policy"`

	// 安全组规则优先级。取值范围：1~100 默认值：1
	Priority *SecurityGroupPriority `json:"priority,omitempty"`

	// 经典网络类型安全组规则的网卡类型。取值范围：
	// - internet：公网网卡
	// - intranet：内网网卡
	// 	- 专有网络VPC类型安全组规则无需设置网卡类型，默认为intranet，只能为intranet。
	// 	- 设置安全组之间互相访问时，即指定了DestGroupId且没有指定DestCidrIp，只能为intranet。
	//
	// 默认值：internet
	NicType *NicType `json:"nic_type"`

	// 安全组规则的描述信息。长度为1~512个字符。
	Description *string `json:"description"`

	// 源端IPv6 CIDR地址段。支持CIDR格式和IPv6格式的IP地址范围。默认值：无
	//
	// @GSD:NOTE 说明 仅支持VPC类型的IP地址。
	IPv6SourceCidrIP *string `json:"ipv6_source_cidr_ip,omitempty"`

	// 目的端IPv6 CIDR地址段。支持CIDR格式和IPv6格式的IP地址范围。默认值：无
	//
	// @GSD:NOTE 说明 仅支持VPC类型的IP地址。
	IPv6DestCidrIP *string `json:"ipv6_dest_cidr_ip,omitempty"`
}

// CreateEgressRuleResponse 安全组出网规则创建返回数据
type CreateEgressRuleResponse struct {
	RequestID string `json:"request_id"`
}

// CreateEgressRule 安全组出网规则创建
func (s *ECS) CreateEgressRule(p *CreateEgressRuleParams) (resp *CreateEgressRuleResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/sgr/%s/egress", p.SecurityGroupID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
