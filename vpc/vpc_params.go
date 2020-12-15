package vpc

import "time"

// VPCInfo 专有网络
type VPCInfo struct {
	ID              string                    `json:"_id"`
	UID             uint32                    `json:"uid"`
	VpcID           string                    `json:"vpc_id"`
	RegionID        string                    `json:"region_id"`
	Status          VPCPendingAvailableStatus `json:"status"` // enum Pending | Available
	VpcName         string                    `json:"vpc_name"`
	VSwitchIDs      VSwitchIdsType            `json:"v_switch_ids"`
	CidrBlock       string                    `json:"cidr_block"`
	VRouterID       string                    `json:"v_router_id"`
	Description     string                    `json:"description"`
	IsDefault       bool                      `json:"is_default"`        // 是否为默认vpc，只能在创建的时候插入该字段
	IPv6CidrBlock   string                    `json:"ipv6_cidr_block"`   // 交换机的IPv6网段
	BindIPv6Gateway bool                      `json:"bind_ipv6_gateway"` // 是否绑定ipv6网关
	CenStatus       CenStatusType             `json:"cen_status"`
	CreateUserCidr  string                    `json:"user_cidr"`
	UserCidrs       UserCidrsType             `json:"user_cidrs"`
	RouteTableID    string                    `json:"route_table_id"`
	ClientToken     string                    `json:"client_token"`
	UpdatedAt       time.Time                 `json:"updated_at"`
	CreatedAt       time.Time                 `json:"created_at"`
}

type VPCPendingAvailableStatus string //enum Pending | Available

const (
	PendingVPCPendingAvailableStatus   VPCPendingAvailableStatus = "Pending"
	AvailableVPCPendingAvailableStatus VPCPendingAvailableStatus = "Available"
)

// CenStatusType VPC绑定云企业网的状态
type CenStatusType = string

// 绑定和未绑定
const (
	DetachedCenStatus  CenStatusType = "Detached"
	AvailableCenStatus CenStatusType = "Available"
)

type VSwitchStatus string //enum Pending | Available

const (
	PendingVSwitchStatus   VSwitchStatus = "Pending"
	AvailableVSwitchStatus VSwitchStatus = "Available"
)

type VPNStatus string

const (
	InitVPN         = VPNStatus("init")
	ProvisioningVPN = VPNStatus("provisioning")
	ActiveVPN       = VPNStatus("active")
	UpdatingVPN     = VPNStatus("updating")
	DeletingVPN     = VPNStatus("deleting")
)

type VPNBusinessStatus string

const (
	VPNNormalBusiness          = VPNBusinessStatus("Normal")
	VPNFinancialLockedBusiness = VPNBusinessStatus("FinancialLocked")
)

type IpsecVPNStatus string

const (
	IkeSaNotEstablished   = IpsecVPNStatus("ike_sa_not_established")
	IkeSaEstablished      = IpsecVPNStatus("ike_sa_established")
	IpsecSaNotEstablished = IpsecVPNStatus("ipsec_sa_not_established")
	IpsecSaEstablished    = IpsecVPNStatus("ipsec_sa_established")
)

type SSLProto string

const (
	UDPSSLProto = SSLProto("UDP")
	TCPSSLProto = SSLProto("TCP")
)

type Cipher string

const (
	AES128CBC = Cipher("AES-128-CBC")
	AES192CBC = Cipher("AES-192-CBC")
	AES256CBC = Cipher("AES-256-CBC")
	NoneEncry = Cipher("none")
)

type SslVpnClientCertStatus string

const (
	ExpiringSslVpnClientCert      = SslVpnClientCertStatus("expiring-soon")
	NormalSslVpnClientCertStatus  = SslVpnClientCertStatus("normal")
	ExpiredSslVpnClientCertStatus = SslVpnClientCertStatus("expired")
)

type VpnInstanceChargeType string

const (
	VpnPREPAY  = VpnInstanceChargeType("PREPAY")
	VpnPOSTPAY = VpnInstanceChargeType("POSTPAY")
)

type VpnSwitch string

const (
	DisableVpnSwitch = VpnSwitch("disable")
	EnableVpnSwitch  = VpnSwitch("enable")
)

// VSwitchIdsType 交换机id列表
type VSwitchIdsType struct {
	VSwitchId []string `json:"v_switch_id"`
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
