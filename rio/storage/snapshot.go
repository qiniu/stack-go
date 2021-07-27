package storage

import "github.com/qiniu/stack-go/components/client"

// Snapshot 快照client
type Snapshot struct {
	client *client.Client
}

// NewSnapshot 获得Snapshot实例
func NewSnapshot(cli *client.Client) *Snapshot {
	return &Snapshot{client: cli}
}
