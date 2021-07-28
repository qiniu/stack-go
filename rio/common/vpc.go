package common

// VPCStatus vpc状态
type VPCStatus string //enum Pending | Available

// 状态枚举
const (
	PendingVPCStatus   VPCStatus = "Pending"
	AvailableVPCStatus VPCStatus = "Available"
)

// VSwitchStatus 交换机状态
type VSwitchStatus string //enum Pending | Available

// 是否可用
const (
	PendingVSwitchStatus   VSwitchStatus = "Pending"
	AvailableVSwitchStatus VSwitchStatus = "Available"
)
