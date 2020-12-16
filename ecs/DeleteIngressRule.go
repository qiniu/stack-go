package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteIngressRuleParams 安全组入网规则删除参数
type DeleteIngressRuleParams struct {
	SecurityGroupID string `json:"security_group_id"`
	RegionID        string `json:"region_id"`

	SecurityGroupRuleID string `json:"security_group_rule_id"`

	// 传输层协议。参数值大小写敏感。取值范围：
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

	// 需要撤销访问权限的源端安全组ID。
	// - 必须设置SourceGroupId或者SourceCidrIp参数。
	// - 如果指定了SourceGroupId没有指定参数SourceCidrIp，则参数NicType取值只能为intranet。
	// - 如果同时指定了SourceGroupId和SourceCidrIp，则默认以SourceCidrIp为准。
	SourceGroupID *string `json:"source_group_id"`

	// 跨账户删除安全组规则时，源端安全组所属的七牛云账户。
	// - 如果SourceGroupOwnerAccount及SourceGroupOwnerId均未设置，则认为是撤销您其他安全组的访问权限。
	// - 如果已经设置参数SourceCidrIp，则参数SourceGroupOwnerAccount无效。
	SourceGroupOwnerAccount *string `json:"source_group_owner_account"`

	// 跨账户删除安全组规则时，源端安全组所属的七牛云账户ID。
	// - 如果SourceGroupOwnerId及SourceGroupOwnerAccount均未设置，则认为是撤销您其他安全组的访问权限。
	// - 如果您已经设置参数SourceCidrIp，则参数SourceGroupOwnerId无效。
	SourceGroupOwnerID *string `json:"source_group_owner_id"`

	// 源端安全组开放的传输层协议相关的端口范围。取值范围：
	//
	// - TCP/UDP协议：取值范围为1~65535。使用斜线（/）隔开起始端口和终止端口。例如：1/200
	// - ICMP协议：-1/-1
	// - GRE协议：-1/-1
	// - all：-1/-1
	SourcePortRange *string `json:"source_port_range"`

	// 源端IP地址范围。支持CIDR格式和IPv4格式的IP地址范围。默认值：0.0.0.0/0
	SourceCidrIP *string `json:"source_cidr_ip"` // IPv4 only, default 0.0.0.0/0

	// 目的端IP地址范围。支持CIDR格式和IPv4格式的IP地址范围。默认值：0.0.0.0/0
	DestCidrIP *string `json:"dest_cidr_ip"` // IPv4 only

	// 访问权限。取值范围：
	// - accept：接受访问
	// - drop：拒绝访问，不发回拒绝信息
	//
	// 默认值：accept
	Policy *PermissionPolicy `json:"policy"` // enum of accept (default) | drop

	// 经典网络类型安全组规则的网卡类型。取值范围：
	// - internet：公网网卡
	// - intranet：内网网卡
	//
	// 默认值：internet
	//
	// 在以下情况中，参数NicType取值只能为intranet：
	// - 专有网络VPC类型安全组规则无需设置网卡类型，默认为intranet，只能为intranet。
	// - 设置安全组之间互相访问时，即指定了DestGroupId且没有指定DestCidrIp，只能为intranet。
	NicType *NicType `json:"nic_type"` // enum of internet | intranet (default)

	// 安全组规则优先级。取值范围：1~100 默认值：1
	Priority *SecurityGroupPriority `json:"priority,omitempty,"` // 1 - 100, default 1 //enum

	// 源端IPv6 CIDR地址段。支持CIDR格式和IPv6格式的IP地址范围。 默认值：无
	//
	// @GSD:NOTE 说明 仅支持VPC类型的IP地址。
	IPv6SourceCidrIP *string `json:"ipv6_source_cidr_ip,omitempty"`

	// 目的端IPv6 CIDR地址段。支持CIDR格式和IPv6格式的IP地址范围。 默认值：无
	//
	// @GSD:NOTE 说明 仅支持VPC类型的IP地址。
	IPv6DestCidrIP *string `json:"ipv6_dest_cidr_ip,omitempty"`
}

// DeleteIngressRuleResponse 安全组入网规则删除返回数据
type DeleteIngressRuleResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteIngressRule 安全组入网规则删除
func (s *ECS) DeleteIngressRule(p *DeleteIngressRuleParams) (resp *DeleteIngressRuleResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/sgr/%s/ingress", p.SecurityGroupID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
