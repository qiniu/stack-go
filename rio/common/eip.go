package common

//EIPStatus eip状态
type EIPStatus string

//const 状态常量
const (
	CreatingEIPStatus      EIPStatus = "Creating"
	AssociatingEIPStatus   EIPStatus = "Associating"
	UnassociatingEIPStatus EIPStatus = "Unassociating"
	InUseEIPStatus         EIPStatus = "InUse"
	AvailableEIPStatus     EIPStatus = "Available"
)

//EIPISP 线路类型
type EIPISP string

// 类型枚举
const (
	BGPProEIPISP EIPISP = "BGP_PRO"
	BGPEIPISP    EIPISP = "BGP"
)
