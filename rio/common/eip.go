package common

//EIPStatus eip状态
type EIPStatus string

//const 状态常量
const (
	CreatingEIPStatus      EIPStatus = "Creating"
	AssociatingEIPStatus   EIPStatus = "Associating"   // 绑定中
	UnassociatingEIPStatus EIPStatus = "Unassociating" // 解绑中
	InUseEIPStatus         EIPStatus = "InUse"
	AvailableEIPStatus     EIPStatus = "Available"
)

//EIPISP 线路类型
type EIPISP string

// 类型枚举
const (
	// BGPProEIPISP BGP（多线）精品线路
	BGPProEIPISP EIPISP = "BGP_PRO"
	// BGPEIPISP BGP（多线）线路
	BGPEIPISP EIPISP = "BGP"
)
