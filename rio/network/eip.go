package network

import "github.com/qiniu/stack-go/components/client"

//EIP 接口封装
type EIP struct {
	client *client.Client
}

//EIPInfo 封装
type EIPInfo struct {
	ZoneID     string `json:"zone_id"`
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
	EIPID      string `json:"eip_id"`
	EIPName    string `json:"eip_name"`
	EIPAddress string `json:"eip_address"`
	Bandwidth  uint   `json:"bandwidth"`
	CreatedAt  int64  `json:"created_at"`
}

//NewEIP 获得Eip实例
func NewEIP(cli *client.Client) *EIP {
	return &EIP{client: cli}
}

// LockReason 锁
type LockReason string //enum financial | security
// LockReason 常量
const (
	FinancialLockReason LockReason = "financial"
	SecurityLockReason  LockReason = "security"
)

// EIPStatus 状态
type EIPStatus string // enum Associating | Unassociating | InUse | Available
// EIPStatus 常量
const (
	CreatingEIPStatus      EIPStatus = "Creating"
	AssociatingEIPStatus   EIPStatus = "Associating"
	UnassociatingEIPStatus EIPStatus = "Unassociating"
	InUseEIPStatus         EIPStatus = "InUse"
	AvailableEIPStatus     EIPStatus = "Available"
)

// EIPISP 线路类型
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
