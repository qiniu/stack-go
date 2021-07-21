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
