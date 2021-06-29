package common

// SourceDiskType 源磁盘类型
type SourceDiskType string

// https://help.aliyun.com/document_detail/25526.html
const (
	SystemSourceDiskType SourceDiskType = "system"
	DataSourceDiskType   SourceDiskType = "data"
)

type UsageType string

const (
	ImageUsageType     UsageType = "image"
	DiskUsageType      UsageType = "disk"
	ImageDiskUsageType UsageType = "image_disk"
	NoneUsageType      UsageType = "none"
)

type SnapshotStatus string

const (
	ProgressingSnapshotStatus  SnapshotStatus = "progressing"
	AccomplishedSnapshotStatus SnapshotStatus = "accomplished"
	FailedSnapshotStatus       SnapshotStatus = "failed"
)

type SnapshotType string

const (
	AutoSnapshotType SnapshotType = "auto"
	UserSnapshotType SnapshotType = "user"
	AllSnapshotType  SnapshotType = "all"
)
