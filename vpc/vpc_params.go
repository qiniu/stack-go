package vpc

import (
	"time"

	"github.com/qiniu/stack-go/params"
)

// VPCInfo 专有网络
type VPCInfo struct {
	ID              string         `json:"_id"`
	UID             uint32         `json:"uid"`
	VpcID           string         `json:"vpc_id"`
	RegionID        string         `json:"region_id"`
	Status          VPCStatus      `json:"status"` // enum Pending | Available
	VpcName         string         `json:"vpc_name"`
	VSwitchIDs      VSwitchIdsType `json:"v_switch_ids"`
	CidrBlock       string         `json:"cidr_block"`
	VRouterID       string         `json:"v_router_id"`
	Description     string         `json:"description"`
	IsDefault       bool           `json:"is_default"`        // 是否为默认vpc，只能在创建的时候插入该字段
	IPv6CidrBlock   string         `json:"ipv6_cidr_block"`   // 交换机的IPv6网段
	BindIPv6Gateway bool           `json:"bind_ipv6_gateway"` // 是否绑定ipv6网关
	CenStatus       CenStatusType  `json:"cen_status"`
	CreateUserCidr  string         `json:"user_cidr"`
	UserCidrs       UserCidrsType  `json:"user_cidrs"`
	RouteTableID    string         `json:"route_table_id"`
	ClientToken     string         `json:"client_token"`
	UpdatedAt       time.Time      `json:"updated_at"`
	CreatedAt       time.Time      `json:"created_at"`
}

// VPCStatus VPC状态
type VPCStatus string //enum Pending | Available

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
type VSwitchStatus string //enum Pending | Available

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
	RegionID string `json:"region_id"` // VPC所在的地域。您可以通过调用DescribeRegions接口获取地域ID。

	// VPC的网段。您可以使用以下网段或其子集作为VPC的网段：
	// - 172.16.0.0/12（默认值）。
	// - 10.0.0.0/8。
	// - 192.168.0.0/16。
	CidrBlock *string `json:"cidr_block"`

	// VPC的名称。
	//
	// 长度为2~128个字符，必须以字母或中文开头，可包含数字、点号（.）、下划线（_）和短横线（-），但不能以 `http://` 或 `https://` 开头。
	VpcName *string `json:"vpc_name"`

	// VPC的描述信息。
	//
	// 长度为2~256个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	Description *string `json:"description"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`

	// 用户网段，如需定义多个网段请使用半角逗号隔开，最多支持3个网段
	UserCidr *string `json:"user_cidr"`

	// VPC的IPv6网段，目前只支持自动分配
	IPv6CidrBlock *string `json:"ipv6_cidr_block,omitempty"`

	// 是否开启IPv6网段，取值：false（默认值）：不开启, true：开启。
	EnableIpv6 *bool `json:"enable_ipv6,omitempty"`
}

// CreateVSwitch 创建交换机
type CreateVSwitch struct {
	// 要创建的交换机所属的可用区ID。您可以通过调用DescribeZones接口获取可用区ID。
	ZoneID string `json:"zone_id"`

	// 要创建的交换机的地域ID。您可以通过调用DescribeRegions接口获取地域ID。
	RegionID string `json:"region_id"`

	// 交换机的网段。交换机网段要求如下：
	// - 交换机的网段的掩码长度范围为16~29位。
	// - 交换机的网段必须从属于所在VPC的网段。
	// - 交换机的网段不能与所在VPC中路由条目的目标网段相同，但可以是目标网段的子集。
	CidrBlock string `json:"cidr_block"`

	// 要创建的交换机所属的VPC ID。
	VPCID string `json:"vpc_id"`

	// 交换机的名称。 名称长度为2~128个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	VSwitchName *string `json:"v_switch_name"`

	// 交换机的描述信息。 描述长度为2~256个字符，必须以字母或中文开头，但不能以 `http://` 或 `https://` 开头。
	Description *string `json:"description"`

	// 交换机IPv6网段的最后8比特位，取值：0~255。
	// 重设json的key，避免和model内的ipv6_cidr_block类型不同convert失败
	IPv6CidrBlock *int `json:"ipv6_cidr_block_number"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`
}

// VSwitchInfo 交换机
type VSwitchInfo struct {
	ID                      string        `json:"-"`
	UID                     uint32        `json:"uid"`
	VSwitchID               string        `json:"v_switch_id"`
	VPCID                   string        `json:"vpc_id"`
	Status                  VSwitchStatus `json:"status"` // enum Pending | Available
	CidrBlock               string        `json:"cidr_block"`
	ZoneID                  string        `json:"zone_id"`
	RegionID                string        `json:"region_id"`
	AvailableIPAddressCount int           `json:"available_ip_address_count"`
	Description             string        `json:"description"`
	VSwitchName             string        `json:"v_switch_name"`
	IsDefault               bool          `json:"is_default"` // 是否为默认交换机，只能在创建的时候插入该字段
	IPv6CidrBlock           string        `json:"ipv6_cidr_block"` // 交换机的IPv6网段
	UpdatedAt               time.Time     `json:"updated_at"`
	CreatedAt               time.Time     `json:"created_at"`
}

// EIPInfo eip
type EIPInfo struct {
	ID             string    `json:"_id"`
	UID            uint32    `json:"uid"`
	AllocationID   string    `json:"allocation_id"`
	AllocationName string    `json:"allocation_name"`
	ChargeType     string    `json:"charge_type"` // enum PrePaid | PostPaid    defulat-PostPaid
	Status         EIPStatus `json:"status"`      // enum Associating | Unassociating | InUse | Available
	InstanceID     string    `json:"instance_id"`
	InstanceName   string    `json:"instance_name"`
	RegionID       string    `json:"region_id"`
	IPAddress      string    `json:"ip_address"`
	ISP            EIPISP    `json:"isp"`

	InternetChargeType string       `json:"internet_charge_type"` // enum PayByBandwidth|PayByTraffic  default-PayByBandwidth
	InstanceType       InstanceType `json:"instance_type"`        // enum EcsInstance | SlbInstance | Nat | HaVip
	Bandwidth          string       `json:"bandwidth"`            // 带宽值，如果加入带宽包，则为被分配的带宽大小
	PricingCycle       string       `json:"pricing_cycle"`
	Period             string       `json:"period"`

	// 带宽包信息，目前只有共享带宽
	BandwidthPackageBandwidth string        `json:"bandwidth_package_bandwidth"` // 带宽包大小
	BandwidthPackageID        string        `json:"bandwidth_package_id"`        // 带宽包id
	BandwidthPackageType      BandwidthType `json:"bandwidth_package_type"`      // 带宽包类型
	EipBandwidth              string        `json:"eip_bandwidth"`               // 当eip加入共享带宽时，该值为加入前的带宽值

	ExpiredTime time.Time `json:"expired_time"`
	UpdatedAt   time.Time `json:"updated_at"`
	CreatedAt   time.Time `json:"created_at"`

	// autoRenewInfo
	AutoRenewInfo        *params.AutoRenew `json:"auto_renew_info"`
	CustomizedExpireTime *time.Time        `json:"customized_expire_time"`

	*params.CostInfo
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
