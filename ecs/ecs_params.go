package ecs

import (
	"time"

	"github.com/qiniu/stack-go/params"
)

// Instance 主机
type Instance struct {
	*params.CostInfo
	ID                 string            `json:"_id"`
	UID                uint32            `json:"uid"`
	InstanceID         string            `json:"instance_id"`
	ImageInfo          ImageInfo         `json:"image_info"`
	ZoneID             string            `json:"zone_id"`
	IoOptimized        bool              `json:"io_optimized"`
	Memory             int64             `json:"memory"`
	CPU                int64             `json:"cpu"`
	InstanceName       string            `json:"instance_name"`
	Description        string            `json:"description"`
	OSType             OsType            `json:"os_type"`
	OSName             string            `json:"os_name"`
	HostName           string            `json:"host_name"`
	InstanceType       string            `json:"instance_type"`
	InstanceTypeFamily string            `json:"instance_type_family"`
	Status             InstanceStatus    `json:"status"`
	RegionID           string            `json:"region_id"`
	GPUSpec            string            `json:"gpu_spec"`
	GPUAmount          int64             `json:"gpu_amount"`
	VncPasswd          string            `json:"vnc_passwd"`
	KeyPairName        string            `json:"key_pair_name"`
	ExpiredTime        time.Time         `json:"expired_time"`
	VpcAttributes      VpcAttributes     `json:"vpc_attributes"`
	SecurityGroupIds   SecurityGroupIds  `json:"security_group_ids"`
	EipAddress         EipAddress        `json:"eip_address"`
	HumanSpec          string            `json:"human_spec"`
	EnableAssignIPv6   bool              `json:"enable_assign_ipv6"`
	Tags               []*params.Tag     `json:"tags"`
	AutoRenewInfo      *params.AutoRenew `json:"auto_renew_info"`
	UpdatedAt          time.Time         `json:"updated_at"`
	CreatedAt          time.Time         `json:"created_at"`
}

// ImageInfo 镜像信息
type ImageInfo struct {
	*params.CostParams

	ImageID              string              `json:"image_id"`
	Description          string              `json:"description"`
	ProductCode          string              `json:"product_code"`
	OSType               string              `json:"os_type"`
	Architecture         string              `json:"architecture"`
	OSName               string              `json:"os_name"`
	ImageOwnerAlias      string              `json:"image_owner_alias"`
	Progress             string              `json:"progress"`
	Usage                string              `json:"usage"`
	ImageVersion         string              `json:"image_version"`
	Status               ImageStatus         `json:"status"`
	ImageName            string              `json:"image_name"`
	IsSelfShared         string              `json:"is_self_shared"`
	Platform             string              `json:"platform"`
	Size                 int64               `json:"size"`
	IsSupportCloudinit   bool                `json:"is_support_cloudinit"`
	IsSupportIoOptimized bool                `json:"is_support_io_optimized"`
	IsCopied             bool                `json:"is_copied"`
	IsSubscribed         bool                `json:"is_subscribed"`
	DiskDeviceMappings   *DiskDeviceMappings `json:"disk_device_mappings"`
	SnapshotID           string              `json:"snapshot_id"`
	MinDiskSize          string              `json:"min_disk_size"`
}

// OsType 操作系统类型
type OsType string

// 操作系统类型
const (
	WindowsOsType OsType = "windows"
	LinuxOsType   OsType = "linux"
)

// InstanceStatus 主机状态
type InstanceStatus string

// 主机状态
const (
	NormalInstanceStatus    InstanceStatus = "Normal"
	CreatingInstanceStatus  InstanceStatus = "Creating"
	ChangingInstanceStatus  InstanceStatus = "Changing"
	RunningInstanceStatus   InstanceStatus = "Running"
	StartingInstanceStatus  InstanceStatus = "Starting"
	StoppingInstanceStatus  InstanceStatus = "Stopping"
	StoppedInstanceStatus   InstanceStatus = "Stopped"
	InactiveInstanceStatus  InstanceStatus = "Inactive"
	RebootingInstanceStatus InstanceStatus = "Rebooting"
	PendingInstanceStatus   InstanceStatus = "Pending"
	DeletedInstanceStatus   InstanceStatus = "Deleted"
)

// VpcAttributes 专有网络参数
type VpcAttributes struct {
	NatIPAddress     string           `json:"nat_ip_address"`
	PrivateIPAddress PrivateIPAddress `json:"private_ip_address"`
	VSwitchID        string           `json:"v_switch_id"`
	VpcID            string           `json:"vpc_id"`
	VpcName          string           `json:"vpc_name"`
	VSwitchName      string           `json:"v_switch_name"`
}

// PrivateIPAddress 内网地址
type PrivateIPAddress struct {
	IPAddress []interface{} `json:"ip_address"`
}

// SecurityGroupIds 安全组
type SecurityGroupIds struct {
	SecurityGroupID   []string `json:"security_group_id"`
	SecurityGroupName []string `json:"security_group_name"`
}

// EipAddress 弹性公网IP
type EipAddress struct {
	IPAddress            string `json:"ip_address"`
	AllocationID         string `json:"allocation_id"`
	InternetChargeType   string `json:"internet_charge_type"`
	IsSupportUnassociate bool   `json:"is_support_unassociate"`
	Bandwidth            int64  `json:"bandwidth"`
}

// ImageStatus 镜像状态
type ImageStatus string

// 镜像状态
const (
	UnAvailableImageStatus  ImageStatus = "UnAvailable"
	AvailableImageStatus    ImageStatus = "Available"
	CreatingImageStatus     ImageStatus = "Creating"
	CreateFailedImageStatus ImageStatus = "CreateFailed"
)

// DiskDeviceMappings 磁盘
type DiskDeviceMappings struct {
	DiskDeviceMapping []*DiskDeviceMapping `json:"disk_device_mapping"`
}

// DiskDeviceMapping 磁盘
type DiskDeviceMapping struct {
	ImportOSSObject string `json:"import_oss_object"`
	Format          string `json:"format"`
	Device          string `json:"device"`
	Type            string `json:"type"`
	SnapshotID      string `json:"snapshot_id"`
	SnapshotName    string `json:"snapshot_name"`
	ImportOSSBucket string `json:"import_oss_bucket"`
	Size            string `json:"size"`
}

// DiskCategory 云盘种类
type DiskCategory string

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

	// DiskCategoryHWESSD 极速型 SSD
	DiskCategoryHWESSD DiskCategory = "H_ESSD"
	// DiskCategoryHWSSD 超高 IO
	DiskCategoryHWSSD DiskCategory = "H_SSD"
	// DiskCategoryHWGPSSD 通用型 SSD
	DiskCategoryHWGPSSD DiskCategory = "H_GPSSD"
	// DiskCategoryHWSAS 高 IO
	DiskCategoryHWSAS DiskCategory = "H_SAS"
	// DiskCategoryHWSata 普通 IO
	DiskCategoryHWSata DiskCategory = "H_SATA"
)

type NetworkType string

const (
	ClassicNetworkType  NetworkType = "Classic"
	VpcNetworkType      NetworkType = "Vpc"
	RedisVpcNetworkType NetworkType = "VPC" // redis 和 rds 的网络类型都使用该参数
)

type AvailableZoneType struct {
	RegionId           string              `json:"region_id"`                     // 地域ID。
	ZoneId             string              `json:"zone_id"`                       // 可用区ID。
	Status             ResourceStatusType  `json:"status"`                        // 资源状态。可能值：`Available`：资源充足 `SoldOut`：资源已售罄
	AvailableResources *AvailableResources `json:"available_resources,omitempty"` // 可供创建的具体资源组成的数组。
}

type AvailableResources struct {
	AvailableResource []*AvailableResource `json:"available_resources,omitempty"`
}

type AvailableResource struct {
	Type               params.ResourceType `json:"type"`
	SupportedResources *SupportedResources `json:"supported_resources,omitempty"` // 支持的可供创建的具体资源组成的数组。
}

type SupportedResources struct {
	SupportedResource []*SupportedResource `json:"supported_resource,omitempty"`
}

type SupportedResource struct {
	Value  string             `json:"value"`          // 资源值。
	Status ResourceStatusType `json:"status"`         // 资源状态。可能值：`Available`：资源充足 `SoldOut`：资源已售罄
	Min    *int64             `json:"min,omitempty"`  // 资源规格的最小值，该参数值为空时将不返回。
	Max    *int64             `json:"max,omitempty"`  // 资源规格的最大值，该参数值为空时将不返回。
	Unit   *string            `json:"unit,omitempty"` // 资源规格单位，该参数值为空时将不返回。
}

type ResourceStatusType string

const (
	AvailableResourceStatusType ResourceStatusType = "Available"
	SoldOutResourceStatusType   ResourceStatusType = "SoldOut"
)
