package storage

import "github.com/qiniu/stack-go/rio/common"

//SnapshotInfo 快照信息
type SnapshotInfo struct {
	SnapshotID   string                `json:"snapshot_id"`   // 快照ID
	SnapshotName string                `json:"snapshot_name"` // 快照显示名称。如果创建时指定了快照显示名称，则返回。
	Description  string                `json:"description"`   // 描述信息
	Status       common.SnapshotStatus `json:"status"`        // 快照状态。可能值： `progressing`, `accomplished`, `failed`
	Usage        common.UsageType      `json:"usage"`         // 快照是否被用作创建镜像或云盘。可能值：`image`, `disk`, `image_disk`, `none`
	Progress     string                `json:"progress"`      // 快照创建进度，单位为百分比。
	RemainTime   int64                 `json:"remain_time"`   // 正在创建的快照剩余完成时间，单位为秒
	SourceDisk   SnapshotDiskInfo      `json:"source_disk"`
	CreatedAt    int64                 `json:"created_at"` // 创建时间。unix 时间戳 ms
}

//SnapshotDiskInfo .
type SnapshotDiskInfo struct {
	DiskID   string          `json:"disk_id"`
	DiskName string          `json:"disk_name"`
	ZoneID   string          `json:"zone_id"`
	DiskType common.DiskType `json:"disk_type"`
	Size     int64           `json:"size"`
}
