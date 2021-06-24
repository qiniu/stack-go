package storage

import "github.com/qiniu/stack-go/components/client"

// Disk 磁盘类接口封装
type Disk struct {
	client *client.Client
}

// NewDisk 获得 Disk 实例
func NewDisk(cli *client.Client) *Disk {
	return &Disk{client: cli}
}
