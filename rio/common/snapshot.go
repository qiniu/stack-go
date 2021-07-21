package common

// SourceDiskType 源磁盘类型
type SourceDiskType string

// https://help.aliyun.com/document_detail/25526.html
const (
	SystemSourceDiskType SourceDiskType = "system"
	DataSourceDiskType   SourceDiskType = "data"
)

// UsageType usage类型
type UsageType string

// const UsageType
const (
	ImageUsageType     UsageType = "image"
	DiskUsageType      UsageType = "disk"
	ImageDiskUsageType UsageType = "image_disk"
	NoneUsageType      UsageType = "none"
)

// SnapshotStatus 快照状态
type SnapshotStatus string

// const SnapshotStatus
const (
	ProgressingSnapshotStatus  SnapshotStatus = "progressing"
	AccomplishedSnapshotStatus SnapshotStatus = "accomplished"
	FailedSnapshotStatus       SnapshotStatus = "failed"
)

// SnapshotType 快照类型
type SnapshotType string

// const SnapshotType
const (
	AutoSnapshotType SnapshotType = "auto"
	UserSnapshotType SnapshotType = "user"
	AllSnapshotType  SnapshotType = "all"
)