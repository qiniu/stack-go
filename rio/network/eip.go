package network

import "github.com/qiniu/stack-go/components/client"

// EIP eip client
type EIP struct {
	client *client.Client
}

// NewEIP 获得Eip实例
func NewEIP(cli *client.Client) *EIP {
	return &EIP{client: cli}
}
