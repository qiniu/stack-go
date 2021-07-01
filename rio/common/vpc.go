package common

type VPCStatus string //enum Pending | Available

const (
	PendingVPCStatus   VPCStatus = "Pending"
	AvailableVPCStatus VPCStatus = "Available"
)

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
