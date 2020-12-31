package vpc

import (
	"time"

	"github.com/qiniu/stack-go/params"
)

// VPCInfo 专有网络
type VPCInfo struct {
	VpcID           string         `json:"vpc_id"`            // VPC ID
	RegionID        string         `json:"region_id"`         // VPC所在的地域
	Status          VPCStatus      `json:"status"`            // VPC的状态，取值：Pending：配置中， Available：可用
	VpcName         string         `json:"vpc_name"`          // VPC 的名称
	VSwitchIDs      VSwitchIdsType `json:"v_switch_ids"`      // VPC 中交换机的列表
	CidrBlock       string         `json:"cidr_block"`        // VPC 的 IPv4 网段
	VRouterID       string         `json:"v_router_id"`       // VPC 路由器的 ID
	Description     string         `json:"description"`       // VPC 的描述信息
	IsDefault       bool           `json:"is_default"`        // 是否为默认 VPC，只能在创建的时候插入该字段
	IPv6CidrBlock   string         `json:"ipv6_cidr_block"`   // 交换机的 IPv6 网段
	BindIPv6Gateway bool           `json:"bind_ipv6_gateway"` // 是否绑定 IPv6 网关
	CenStatus       CenStatusType  `json:"cen_status"`        // VPC绑定云企业网的状态: Detached：未绑定云企业网， Available：已绑定云企业网
	RouteTableID    string         `json:"route_table_id"`    // 路由表的ID
	ClientToken     string         `json:"client_token"`      // 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一
	CreatedAt       time.Time      `json:"created_at"`        // 创建时间
	UpdatedAt       time.Time      `json:"updated_at"`        // 更新时间
}

// VPCStatus VPC状态
type VPCStatus string

// VPC 状态
const (
	PendingVPCPendingAvailableStatus   VPCStatus = "Pending"
	AvailableVPCPendingAvailableStatus VPCStatus = "Available"
)

// CenStatusType VPC绑定云企业网的状态
type CenStatusType = string

// 绑定和未绑定
const (
	DetachedCenStatus  CenStatusType = "Detached"
	AvailableCenStatus CenStatusType = "Available"
)

// VSwitchStatus VSwitch 状态
type VSwitchStatus string

// VSwitchStatus VSwitch 状态
const (
	PendingVSwitchStatus   VSwitchStatus = "Pending"
	AvailableVSwitchStatus VSwitchStatus = "Available"
)

// VSwitchIdsType 交换机id列表
type VSwitchIdsType struct {
	VSwitchID []string `json:"v_switch_id"`
}

// UserCidrsType 用户侧网段的列表
type UserCidrsType struct {
	UserCidr []string `json:"user_cidr"`
}

// CreateVpc 创建专有网络
type CreateVpc struct {

	// VPC所在的地域
	//
	// 您可以通过调用DescribeRegions接口获取地域ID。
	RegionID string `json:"region_id"`

	// VPC 的网段
	//
	// 您可以使用以下网段或其子集作为VPC的网段：
	// - 172.16.0.0/12（默认值）。
	// - 10.0.0.0/8。
	// - 192.168.0.0/16。
	CidrBlock *string `json:"cidr_block"`

	// VPC 的名称
	//
	// 长度为2~128个字符，必须以字母或中文开头，可包含数字、点号（.）、下划线（_）和短横线（-），但不能以 `http://` 或 `https://` 开头。
	VpcName *string `json:"vpc_name"`

	// VPC的描述信息
	//
	// 长度为2~256个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	Description *string `json:"description"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`

	// 用户网段，如需定义多个网段请使用半角逗号隔开，最多支持3个网段
	UserCidr *string `json:"user_cidr"`

	// VPC的IPv6网段，目前只支持自动分配
	IPv6CidrBlock *string `json:"ipv6_cidr_block,omitempty"`

	// 是否开启IPv6网段
	//
	// 取值：
	// - `false`：不开启（默认值）
	// - `true`：开启。
	EnableIpv6 *bool `json:"enable_ipv6,omitempty"`
}

// CreateVSwitch 创建交换机
type CreateVSwitch struct {

	// 要创建的交换机所属的可用区ID
	//
	// 您可以通过调用DescribeZones接口获取可用区ID。
	ZoneID string `json:"zone_id"`

	// 要创建的交换机的地域ID
	//
	// 您可以通过调用DescribeRegions接口获取地域ID。
	RegionID string `json:"region_id"`

	// 交换机的网段。交换机网段要求如下：
	// - 交换机的网段的掩码长度范围为16~29位。
	// - 交换机的网段必须从属于所在VPC的网段。
	// - 交换机的网段不能与所在VPC中路由条目的目标网段相同，但可以是目标网段的子集。
	CidrBlock string `json:"cidr_block"`

	// 要创建的交换机所属的VPC ID。
	VPCID string `json:"vpc_id"`

	// 交换机的名称
	//
	// 名称长度为2~128个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	VSwitchName *string `json:"v_switch_name"`

	// 交换机的描述信息
	//
	// 描述长度为2~256个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	Description *string `json:"description"`

	// 交换机IPv6网段的最后8比特位
	//
	// 取值：0~255。
	//
	// 重设json的key，避免和model内的ipv6_cidr_block类型不同convert失败
	IPv6CidrBlock *int `json:"ipv6_cidr_block_number"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`
}

// VSwitchInfo 交换机
type VSwitchInfo struct {
	VSwitchID               string        `json:"v_switch_id"`                // 交换机的ID
	VPCID                   string        `json:"vpc_id"`                     // 交换机所属VPC的ID
	Status                  VSwitchStatus `json:"status"`                     // 交换机的状态，取值：Pending：配置中， Available：可用。
	CidrBlock               string        `json:"cidr_block"`                 // 交换机的IPv4网段。
	ZoneID                  string        `json:"zone_id"`                    // 交换机所属的可用区 ID
	RegionID                string        `json:"region_id"`                  // 交换机所属的区域 ID
	AvailableIPAddressCount int           `json:"available_ip_address_count"` // 交换机中可用的IP地址数量
	Description             string        `json:"description"`                // 交换机的描述信息
	VSwitchName             string        `json:"v_switch_name"`              // 交换机的名称
	IsDefault               bool          `json:"is_default"`                 // 是否为默认交换机，只能在创建的时候插入该字段
	IPv6CidrBlock           string        `json:"ipv6_cidr_block"`            // 交换机的IPv6网段
	CreatedAt               time.Time     `json:"created_at"`                 // 创建时间
	UpdatedAt               time.Time     `json:"updated_at"`                 // 更新时间
}

// EIPInfo eip
type EIPInfo struct {
	AllocationID   string `json:"allocation_id"`   // 弹性公网 IP 的 ID
	AllocationName string `json:"allocation_name"` // 弹性公网 IP 的名称

	// 弹性公网IP的计费模式。
	// - PrePaid：包年包月。
	// - PostPaid：按量计费。
	ChargeType string `json:"charge_type"`

	// EIP的状态。
	// - Associating：绑定中。
	// - Unassociating：解绑中。
	// - InUse：已分配。
	// - Available：可用。
	Status EIPStatus `json:"status"`

	// 当前绑定的实例的ID
	InstanceID string `json:"instance_id"`

	// 当前绑定的实例的名称
	InstanceName string `json:"instance_name"`

	// EIP所在的地域
	RegionID string `json:"region_id"`

	// 弹性公网IP的IP地址
	IPAddress string `json:"ip_address"`

	// 服务提供商
	ISP EIPISP `json:"isp"`

	// EIP的计量方式，取值：
	// - PayByBandwidth（默认值）：按带宽计费。
	// - PayByTraffic：按流量计费。
	// - 当InstanceChargeType取值为PrePaid时，InternetChargeType必须取值PayByBandwidth。
	// - 当InstanceChargeType取值为PostPaid时，InternetChargeType可取值PayByBandwidth或PayByTraffic。
	InternetChargeType string `json:"internet_charge_type"`

	// EIP 的实例的类型，取值：
	// - Nat：NAT网关。
	// - SlbInstance：负载均衡SLB。
	// - EcsInstance：云服务器ECS。
	// - NetworkInterface：辅助弹性网卡。
	// - HaVip：高可用虚拟IP。
	InstanceType InstanceType `json:"instance_type"`

	// EIP的带宽峰值，单位为Mbps
	//
	// 如果加入带宽包，则为被分配的带宽大小
	Bandwidth string `json:"bandwidth"`

	// 包年包月的计费周期，取值：
	// - Month（默认值）：按月付费。
	// - Year：按年付费。
	PricingCycle string `json:"pricing_cycle"`

	// 购买时长。
	// - 当PricingCycle取值Month时，Period取值范围为1~9。
	// - 当PricingCycle取值Year时，Period取值范围为1~3。
	Period string `json:"period"`

	// EIP 加入的共享带宽的带宽值
	BandwidthPackageBandwidth string `json:"bandwidth_package_bandwidth"`

	// 加入的共享带宽ID
	BandwidthPackageID string `json:"bandwidth_package_id"`

	// 带宽的类型，仅支持返回CommonBandwidthPackage（共享带宽）
	BandwidthPackageType BandwidthType `json:"bandwidth_package_type"`

	// EIP加入共享带宽之前或退出共享带宽之后的带宽
	EipBandwidth string `json:"eip_bandwidth"`

	ExpiredTime time.Time `json:"expired_time"` // 到期时间
	CreatedAt   time.Time `json:"created_at"`   // 创建时间
	UpdatedAt   time.Time `json:"updated_at"`   // 更新时间

	AutoRenewInfo        *params.AutoRenew `json:"auto_renew_info"`        // 自动续费信息
	CustomizedExpireTime *time.Time        `json:"customized_expire_time"` // 自定义过期时间

	*params.CostInfo // 付费信息
}

// EIPStatus EIP 状态
type EIPStatus string // enum Associating | Unassociating | InUse | Available

// EIPStatus EIP 状态
const (
	CreatingEIPStatus      EIPStatus = "Creating"
	AssociatingEIPStatus   EIPStatus = "Associating"
	UnassociatingEIPStatus EIPStatus = "Unassociating"
	InUseEIPStatus         EIPStatus = "InUse"
	AvailableEIPStatus     EIPStatus = "Available"
)

// EIPISP 线路类型
type EIPISP string

// EIPISP 线路类型
const (
	// 资源池 1
	// BGPProEIPISP BGP（多线）精品线路
	BGPProEIPISP EIPISP = "BGP_PRO"
	// BGPEIPISP BGP（多线）线路
	BGPEIPISP EIPISP = "BGP"

	// 资源池 2
	// TelcomEIPISP2 电信
	TelcomEIPISP2 EIPISP = "5_telcom"
	// UnionEIPISP2 联通
	UnionEIPISP2 EIPISP = "5_union"
	// BGPEIPISP2 全动态BGP
	BGPEIPISP2 EIPISP = "5_bgp"
	// SBGPEIPISP2 静态BGP
	SBGPEIPISP2 EIPISP = "5_sbgp"
)

// InstanceType 实例类型
type InstanceType string // enum EcsInstance | SlbInstance | Nat | HaVip

// InstanceType 实例类型
const (
	EcsInstanceInstanceType InstanceType = "EcsInstance"
	SlbInstanceInstanceType InstanceType = "SlbInstance"
	NatInstanceType         InstanceType = "Nat"
	HaVipInstanceType       InstanceType = "HaVip"
	NullInstanceType        InstanceType = ""
)

// BandwidthType eip带宽类型
type BandwidthType string

// 带宽类型常量
const (
	CommonBandwidthPackage BandwidthType = "CommonBandwidthPackage"
)
