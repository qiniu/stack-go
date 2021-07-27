package common

// DiskCategory 云盘种类
type DiskCategory string

// DiskCategory 云盘种类常量
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

// DiskType 磁盘类型
type DiskType string

// DiskType 云盘类型
const (
	SystemDiskType DiskType = "system"
	DataDiskType   DiskType = "data"
)

// DiskStatusType 磁盘状态类型
type DiskStatusType string

// const 磁盘状态
const (
	DiskInUse         DiskStatusType = "In_use"
	DiskAvailable     DiskStatusType = "Available"
	DiskAttaching     DiskStatusType = "Attaching"
	DiskDetaching     DiskStatusType = "Detaching"
	DiskCreating      DiskStatusType = "Creating"
	DiskInitializing  DiskStatusType = "Initializing"
	DiskReIniting     DiskStatusType = "ReIniting"
	DiskProcessing    DiskStatusType = "Processing"
	AllDiskStatusType DiskStatusType = "All"
)
