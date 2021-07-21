package storage

import "github.com/qiniu/stack-go/components/client"

// Snapshot 快照类接口封装
type Snapshot struct {
	client *client.Client
}

// NewSnapshot 获得Snapshot实例
func NewSnapshot(cli *client.Client) *Snapshot {
	return &Snapshot{client: cli}
}
