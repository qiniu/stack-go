package storage

import "github.com/qiniu/stack-go/components/client"

//Snapshot 快照类接口封装
type Snapshot struct {
	client *client.Client
}

//NewSnapshot 获得Snapshot实例
func NewSnapshot(cli *client.Client) *Snapshot {
	return &Snapshot{client: cli}
}

// SourceDiskType 源磁盘类型
type SourceDiskType string

// SourceDiskType 常量
const (
	SystemSourceDiskType SourceDiskType = "system"
	DataSourceDiskType   SourceDiskType = "data"
)

// UsageType 使用类型
type UsageType string

// UsageType 常量
const (
	ImageUsageType     UsageType = "image"
	DiskUsageType      UsageType = "disk"
	ImageDiskUsageType UsageType = "image_disk"
	NoneUsageType      UsageType = "none"
)

// SnapshotStatus 快照状态
type SnapshotStatus string

// SnapshotStatus 常量
const (
	ProgressingSnapshotStatus  SnapshotStatus = "progressing"
	AccomplishedSnapshotStatus SnapshotStatus = "accomplished"
	FailedSnapshotStatus       SnapshotStatus = "failed"
	RestoringSnapshotStatus    SnapshotStatus = "restoring"
	ProcessingSnapshotStatus   SnapshotStatus = "processing"
)

// SnapshotType 快照类型
type SnapshotType string

// SnapshotType 常量
const (
	AutoSnapshotType SnapshotType = "auto"
	UserSnapshotType SnapshotType = "user"
	AllSnapshotType  SnapshotType = "all"
)
