package ecs

import "github.com/qiniu/stack-go/components/client"

// ECS 主机类接口封装
type ECS struct {
	client *client.Client
}

// NewECS 获得 ECS 实例
func NewECS(cli *client.Client) *ECS {
	return &ECS{
		client: cli,
	}
}
