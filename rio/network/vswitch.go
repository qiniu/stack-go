package network

import (
	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitch 接口封装
type VSwitch struct {
	client           *client.Client
	ZoneID           string               `json:"zone_id"`
	VSwitchID        string               `json:"vswitch_id"`
	VPCID            string               `json:"vpc_id"`
	Status           common.VSwitchStatus `json:"status"`
	CIDR             string               `json:"cidr"`
	AvailableIPCount int                  `json:"available_ip_count"`
	Description      string               `json:"description"`
	VSwitchName      string               `json:"vswitch_name"`
	CreatedAt        int64                `json:"created_at"`
	ResourceCount    int                  `json:"resource_count"`
}

// NewVSwitch 获得VSwitch实例
func NewVSwitch(cli *client.Client) *VSwitch {
	return &VSwitch{client: cli}
}
