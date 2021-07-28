package common

// SecurityRuleGrantType 安全组授权类型
type SecurityRuleGrantType int

// SecurityRuleGrantType 常量
const (
	SecurityGroupType SecurityRuleGrantType = 1 // 安全组授权
	CIDRType          SecurityRuleGrantType = 2 // CIDR 授权
)

// Direction 安全组方向
type Direction string

// 方向常量
const (
	IngressDirection Direction = "ingress" // 进方向
	EgressDirection  Direction = "egress"  // 出方向
	AllDirection     Direction = "all"
)

// NetworkProtocol 网络协议
type NetworkProtocol string

// NetworkProtocol 常量
const (
	AllNetworkProtocol  NetworkProtocol = "all"
	TCPNetworkProtocol  NetworkProtocol = "tcp"
	UDPNetworkProtocol  NetworkProtocol = "udp"
	ICMPNetworkProtocol NetworkProtocol = "icmp"
	GRENetworkProtocol  NetworkProtocol = "gre"
)
