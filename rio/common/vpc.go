package common

// VPCStatus vpc状态
type VPCStatus string //enum Pending | Available
// 状态枚举
const (
	PendingVPCStatus   VPCStatus = "Pending"
	AvailableVPCStatus VPCStatus = "Available"
)

// CenStatusType 绑定类型
type CenStatusType = string

// 绑定和未绑定
const (
	DetachedCenStatus  CenStatusType = "Detached"
	AvailableCenStatus CenStatusType = "Available"
)

// VSwitchStatus 交换机状态
type VSwitchStatus string //enum Pending | Available
// 是否可用
const (
	PendingVSwitchStatus   VSwitchStatus = "Pending"
	AvailableVSwitchStatus VSwitchStatus = "Available"
)

// VPNStatus vpn状态
type VPNStatus string

// 状态枚举
const (
	InitVPN         = VPNStatus("init")
	ProvisioningVPN = VPNStatus("provisioning")
	ActiveVPN       = VPNStatus("active")
	UpdatingVPN     = VPNStatus("updating")
	DeletingVPN     = VPNStatus("deleting")
)

// VPNBusinessStatus 状态
type VPNBusinessStatus string

// 正常和锁
const (
	VPNNormalBusiness          = VPNBusinessStatus("Normal")
	VPNFinancialLockedBusiness = VPNBusinessStatus("FinancialLocked")
)

// IpsecVPNStatus 状态
type IpsecVPNStatus string

// 是否建立
const (
	IkeSaNotEstablished   = IpsecVPNStatus("ike_sa_not_established")
	IkeSaEstablished      = IpsecVPNStatus("ike_sa_established")
	IpsecSaNotEstablished = IpsecVPNStatus("ipsec_sa_not_established")
	IpsecSaEstablished    = IpsecVPNStatus("ipsec_sa_established")
)

// SSLProto tcp/udp
type SSLProto string

// tcp/udp
const (
	UDPSSLProto = SSLProto("UDP")
	TCPSSLProto = SSLProto("TCP")
)

// Cipher 加密
type Cipher string

// 编码
const (
	AES128CBC = Cipher("AES-128-CBC")
	AES192CBC = Cipher("AES-192-CBC")
	AES256CBC = Cipher("AES-256-CBC")
	NoneEncry = Cipher("none")
)

// SslVpnClientCertStatus 状态
type SslVpnClientCertStatus string

// 是否正常
const (
	ExpiringSslVpnClientCert      = SslVpnClientCertStatus("expiring-soon")
	NormalSslVpnClientCertStatus  = SslVpnClientCertStatus("normal")
	ExpiredSslVpnClientCertStatus = SslVpnClientCertStatus("expired")
)

// VpnInstanceChargeType 类型
type VpnInstanceChargeType string

// 支付方式
const (
	VpnPREPAY  = VpnInstanceChargeType("PREPAY")
	VpnPOSTPAY = VpnInstanceChargeType("POSTPAY")
)

// VpnSwitch vpn交换机
type VpnSwitch string

// 是否可用
const (
	DisableVpnSwitch = VpnSwitch("disable")
	EnableVpnSwitch  = VpnSwitch("enable")
)
