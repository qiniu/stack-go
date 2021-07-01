package network

import (
	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//VPC 接口封装
type VPC struct {
	client      *client.Client
	VPCID       string           `json:"vpc_id"`
	ZoneID      string           `json:"zone_id"`
	Status      common.VPCStatus `json:"status"`
	VPCName     string           `json:"vpc_name"`
	VSwitchIDs  []string         `json:"vswitch_ids"`
	CIDR        string           `json:"cidr"`
	Description string           `json:"description"`
	CreatedAt   int64            `json:"created_at"`
}

//NewVPC 获得Vpc实例
func NewVPC(cli *client.Client) *VPC {
	return &VPC{client: cli}
}
