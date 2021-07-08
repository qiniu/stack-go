package compute

import "github.com/qiniu/stack-go/components/client"

// Monitor 镜像类接口封装
type Monitor struct {
	client *client.Client
}

// NewMonitor 获得镜像实例
func NewMonitor(cli *client.Client) *Monitor {
	return &Monitor{client: cli}
}
