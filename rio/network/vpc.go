package network

import (
	"github.com/qiniu/stack-go/components/client"
)

// VPC vpc client
type VPC struct {
	client *client.Client
}

// NewVPC 获得Vpc实例
func NewVPC(cli *client.Client) *VPC {
	return &VPC{client: cli}
}
