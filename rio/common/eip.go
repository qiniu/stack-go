package common

//LockReason .
type LockReason string

//const .
const (
	FinancialLockReason LockReason = "financial"
	SecurityLockReason  LockReason = "security"
)

//EIPStatus eip状态
type EIPStatus string

//const .
const (
	CreatingEIPStatus      EIPStatus = "Creating"
	AssociatingEIPStatus   EIPStatus = "Associating"
	UnassociatingEIPStatus EIPStatus = "Unassociating"
	InUseEIPStatus         EIPStatus = "InUse"
	AvailableEIPStatus     EIPStatus = "Available"
)

//EIPISP 线路类型
type EIPISP string

const (
	// BGPProEIPISP ALI BGP（多线）精品线路
	BGPProEIPISP EIPISP = "BGP_PRO"
	// BGPEIPISP ALI BGP（多线）线路
	BGPEIPISP EIPISP = "BGP"

	// HWTelcomEIPISP HW 电信
	HWTelcomEIPISP EIPISP = "5_telcom"
	// HWUnionEIPISP HW 联通
	HWUnionEIPISP EIPISP = "5_union"
	// HWBGPEIPISP HW 全动态BGP
	HWBGPEIPISP EIPISP = "5_bgp"
	// HWSBGPEIPISP HW 静态BGP
	HWSBGPEIPISP EIPISP = "5_sbgp"
)
