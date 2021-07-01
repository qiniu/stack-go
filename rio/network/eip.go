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
