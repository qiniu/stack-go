package network

import (
	"github.com/qiniu/stack-go/components/client"
)

// VSwitch 交换机
type VSwitch struct {
	client *client.Client
}

// NewVSwitch 获得VSwitch实例
func NewVSwitch(cli *client.Client) *VSwitch {
	return &VSwitch{client: cli}
}
