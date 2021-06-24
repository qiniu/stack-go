package storage

import "github.com/qiniu/stack-go/rio/common"

// DiskInfo 磁盘信息
type DiskInfo struct {
	DiskID           string                `json:"disk_id"`
	DiskName         string                `json:"disk_name"`
	DiskType         common.DiskType       `json:"disk_type"`
	Device           string                `json:"device"`
	Size             int64                 `json:"size"`
	Status           common.DiskStatusType `json:"status"`
	ZoneID           string                `json:"zone_id"`
	SourceSnapshotID string                `json:"source_snapshot_id"`
	Server           DiskServerInfo        `json:"server"`
	ImageID          string                `json:"image_id"`
	Portable         bool                  `json:"portable"`
	Bootable         bool                  `json:"bootable"`
	AttachedTime     int64                 `json:"attached_time"`
	DeleteWithServer bool                  `json:"delete_with_server"`
	Description      string                `json:"description"`
	DetachedAt       int64                 `json:"detached_at"`
	CreatedAt        int64                 `json:"created_at"`
}

// DiskServerInfo 磁盘绑定主机的信息
type DiskServerInfo struct {
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
}
