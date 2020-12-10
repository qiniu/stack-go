package vpc

import "github.com/qiniu/stack-go/components/client"

// VPC 专有网络类接口封装
type VPC struct {
	client *client.Client
}

// NewVPC 获得 VPC 实例
func NewVPC(cli *client.Client) *VPC {
	return &VPC{
		client: cli,
	}
}
