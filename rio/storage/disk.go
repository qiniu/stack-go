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

// DiskCategory 云盘种类
type DiskCategory string

// DiskCategory 常量
const (
	AllDiskCategory             DiskCategory = "all"              // 所有云盘以及本地盘
	CloudDiskCategory           DiskCategory = "cloud"            // 普通云盘
	CloudEfficiencyDiskCategory DiskCategory = "cloud_efficiency" // 高效云盘
	CloudSSDiskCategory         DiskCategory = "cloud_ssd"        // SSD云盘
	CloudESSDiskCategory        DiskCategory = "cloud_essd"       // 增强型SSD云盘

	CloudLocalSSDiskCategory DiskCategory = "cloud_local_ssd" // 本地SSD云盘
	CloudLocalDiskCategory   DiskCategory = "cloud_local"     // 本地云盘
	EphemeralSSDiskCategory  DiskCategory = "ephemeral_ssd"   // 本地SSD盘
	LocalSSDProDiskCategory  DiskCategory = "local_ssd_pro"   // 本地SSD
)

// ESSDPerformanceLevel ESSD云盘性能等级
type ESSDPerformanceLevel string

// ESSDPerformanceLevel 常量
const (
	PL1 ESSDPerformanceLevel = "PL1"
	PL2 ESSDPerformanceLevel = "PL2"
	PL3 ESSDPerformanceLevel = "PL3"
)

// DiskType 磁盘类型
type DiskType string

// DiskType 常量
const (
	AllDiskType    DiskType = "all"
	SystemDiskType DiskType = "system"
	DataDiskType   DiskType = "data"
)

// DiskStatusType 磁盘状态类型
type DiskStatusType string

// DiskStatusType 磁盘状态常量
const (
	DiskInUse         DiskStatusType = "In_use"
	DiskAvailable     DiskStatusType = "Available"
	DiskAttaching     DiskStatusType = "Attaching"
	DiskDetaching     DiskStatusType = "Detaching"
	DiskCreating      DiskStatusType = "Creating"
	DiskInitializing  DiskStatusType = "Initializing"
	DiskReIniting     DiskStatusType = "ReIniting"
	DiskProcessing    DiskStatusType = "Processing"
	DiskDeleting      DiskStatusType = "Deleting"
	DiskError         DiskStatusType = "Error"
	AllDiskStatusType DiskStatusType = "All"
)
