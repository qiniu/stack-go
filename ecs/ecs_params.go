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

// 云盘种类
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

// NetworkType 网络类型
type NetworkType string

// NetworkType 网络类型
const (
	ClassicNetworkType  NetworkType = "Classic"
	VpcNetworkType      NetworkType = "Vpc"
	RedisVpcNetworkType NetworkType = "VPC" // redis 和 rds 的网络类型都使用该参数
)

// AvailableZoneType 可用区信息
type AvailableZoneType struct {
	RegionID           string              `json:"region_id"`                     // 地域ID。
	ZoneID             string              `json:"zone_id"`                       // 可用区ID。
	Status             ResourceStatusType  `json:"status"`                        // 资源状态。可能值：`Available`：资源充足 `SoldOut`：资源已售罄
	AvailableResources *AvailableResources `json:"available_resources,omitempty"` // 可供创建的具体资源组成的数组。
}

// AvailableResources 可用资源信息
type AvailableResources struct {
	AvailableResource []*AvailableResource `json:"available_resources,omitempty"`
}

// AvailableResource 可用资源信息
type AvailableResource struct {
	Type               params.ResourceType `json:"type"`
	SupportedResources *SupportedResources `json:"supported_resources,omitempty"` // 支持的可供创建的具体资源组成的数组。
}

// SupportedResources 支持的资源信息
type SupportedResources struct {
	SupportedResource []*SupportedResource `json:"supported_resource,omitempty"`
}

// SupportedResource 支持的资源信息
type SupportedResource struct {
	Value  string             `json:"value"`          // 资源值。
	Status ResourceStatusType `json:"status"`         // 资源状态。可能值：`Available`：资源充足 `SoldOut`：资源已售罄
	Min    *int64             `json:"min,omitempty"`  // 资源规格的最小值，该参数值为空时将不返回。
	Max    *int64             `json:"max,omitempty"`  // 资源规格的最大值，该参数值为空时将不返回。
	Unit   *string            `json:"unit,omitempty"` // 资源规格单位，该参数值为空时将不返回。
}

// ResourceStatusType 资源状态
type ResourceStatusType string

// ResourceStatusType 资源状态
const (
	AvailableResourceStatusType ResourceStatusType = "Available"
	SoldOutResourceStatusType   ResourceStatusType = "SoldOut"
)

// CreateInstanceEip 创建主机时附带创建EIP信息
type CreateInstanceEip struct {
	EipID        *string `json:"eip_id"`
	EipBandwidth *string `json:"eip_bandwidth"`
	EipName      string  `json:"eip_name"`
	EipAutoPay   string  `json:"eip_auto_pay"`

	*params.CostParams
}

// SecurityEnhancementStrategy 安全加固
type SecurityEnhancementStrategy string

// SecurityEnhancementStrategy 安全加固
const (
	ActiveSecurityEnhancementStrategy   SecurityEnhancementStrategy = "Active"
	DeactiveSecurityEnhancementStrategy SecurityEnhancementStrategy = "Deactive"
)

// DiskInfo 磁盘信息
type DiskInfo struct {
	Size               int                   `json:"size,omitempty"`
	Category           *DiskCategory         `json:"category,omitempty"`
	Encrypted          *bool                 `json:"encrypted,omitempty"`
	SnapshotID         *string               `json:"snapshot_id,omitempty"`
	DiskName           *string               `json:"disk_name,omitempty"`
	Description        *string               `json:"description,omitempty"`
	DeleteWithInstance *bool                 `json:"delete_with_instance,omitempty"`
	PerformanceLevel   *ESSDPerformanceLevel `json:"performance_level,omitempty"`
}

// ESSDPerformanceLevel ESSD云盘性能等级
type ESSDPerformanceLevel string

// ESSDPerformanceLevel ESSD云盘性能等级
const (
	PL1 ESSDPerformanceLevel = "PL1"
	PL2 ESSDPerformanceLevel = "PL2"
	PL3 ESSDPerformanceLevel = "PL3"
)

// InstanceStoppedMode 主机停机状态
type InstanceStoppedMode string

/**
 * KeepChargingInstanceStoppedMode 停机继续收费
 * StopChargingInstanceStoppedMode 停机不收费
 * NotApplicableInstanceStoppedMode 开机
 */
const (
	KeepChargingInstanceStoppedMode  InstanceStoppedMode = "KeepCharging"
	StopChargingInstanceStoppedMode  InstanceStoppedMode = "StopCharging"
	NotApplicableInstanceStoppedMode InstanceStoppedMode = "Not-applicable"
)

// InstanceType 实例类型
type InstanceType struct {
	ID                     string    `json:"_id"`
	NamespaceName          string    `json:"namespace_name"`
	Type                   string    `json:"type"`
	Family                 string    `json:"family"`
	FamilyName             string    `json:"family_name"`
	CPU                    uint      `json:"cpu"`    // 个
	Memory                 float32   `json:"memory"` // GB
	GpuSpec                string    `json:"gpu_spec"`
	GpuAmount              uint      `json:"gpu_amount"`           // 个
	LocalStorageAmount     uint      `json:"local_storage_amount"` // 个
	LocalStorageCategory   string    `json:"local_storage_category"`
	LocalStorageCapacity   float32   `json:"local_storage_capacity"` // GB
	PrivatePPS             float32   `json:"private_pps"`            // 内网收发包 万
	IntranetBandwidth      float32   `json:"intranet_bandwidth"`     // 内网带宽 Gbps
	ClockSpeed             string    `json:"clock_speed"`
	PhysicalProcessor      string    `json:"physical_processor"`
	EniIPv6AddressQuantity uint32    `json:"eni_ipv6_address_quantity"` // 单块弹性网卡的IPv6地址上限
	Architecture           string    `json:"architecture"`
	UpdatedAt              time.Time `json:"updated_at"`
	CreatedAt              time.Time `json:"created_at"`
}

// Disk 磁盘
type Disk struct {
	ID                            string               `json:"_id"`
	UID                           uint32               `json:"uid"`
	ZoneID                        string               `json:"zone_id"`
	RegionID                      string               `json:"region_id"`
	DiskID                        string               `json:"disk_id"`
	DiskName                      string               `json:"disk_name"`
	Type                          DiskType             `json:"type"`
	Size                          int64                `json:"size"`
	Status                        DiskStatusType       `json:"status"`
	Device                        string               `json:"device"`
	IsBind                        bool                 `json:"is_bind"`
	Category                      DiskCategory         `json:"category"`
	Portable                      bool                 `json:"portable"`
	Encrypted                     bool                 `json:"encrypted"`
	ProductCode                   string               `json:"product_code"`
	Description                   string               `json:"description"`
	ExpiredTime                   time.Time            `json:"expired_time"`
	AttachedTime                  time.Time            `json:"attached_time"`
	DetachedTime                  time.Time            `json:"detached_time"`
	DiskChargeType                string               `json:"disk_charge_type"`
	EnableAutoSnapshot            bool                 `json:"enable_auto_snapshot"`
	DeleteAutoSnapshot            bool                 `json:"delete_auto_snapshot"`
	EnableAutomatedSnapshotPolicy bool                 `json:"enable_automated_snapshot_policy"`
	DeleteWithInstance            bool                 `json:"delete_with_instance"`
	ImageID                       string               `json:"image_id"`
	InstanceID                    string               `json:"instance_id"`
	InstanceName                  string               `json:"instance_name"`
	ResourceGroupID               string               `json:"resource_group_id"`
	SourceSnapshotID              string               `json:"source_snapshot_id"`
	AutoSnapshotPolicyID          string               `json:"auto_snapshot_policy_id"`
	PerformanceLevel              ESSDPerformanceLevel `json:"performance_level"`
	IOPS                          int64                `json:"iops"`
	IOPSRead                      int64                `json:"iops_read"`
	IOPSWrite                     int64                `json:"iops_write"`
	CreatedAt                     time.Time            `json:"created_at"`
	UpdatedAt                     time.Time            `json:"updated_at"`

	*params.CostInfo
}

// DiskType 磁盘类型
type DiskType string

// 磁盘类型
const (
	AllDiskType    DiskType = "all"
	SystemDiskType DiskType = "system"
	DataDiskType   DiskType = "data"
)

// DiskStatusType 磁盘状态类型
type DiskStatusType string

// 磁盘状态
const (
	DiskInUse         DiskStatusType = "In_use"
	DiskAvailable     DiskStatusType = "Available"
	DiskAttaching     DiskStatusType = "Attaching"
	DiskDetaching     DiskStatusType = "Detaching"
	DiskCreating      DiskStatusType = "Creating"
	DiskInitializing  DiskStatusType = "Initializing"
	DiskReIniting     DiskStatusType = "ReIniting"
	AllDiskStatusType DiskStatusType = "All"
)

// Keypair 钥匙对
type Keypair struct {
	ID                 string    `json:"id"`
	UID                uint32    `json:"-"`
	KeyPairName        string    `json:"key_pair_name"`
	KeyPairFingerPrint string    `json:"key_pair_finger_print"`
	PrivateKeyBody     string    `json:"private_key_body"`
	PublicKeyBody      string    `json:"public_key_body"`
	CreatedAt          time.Time `json:"created_at"`
}

// ImageOwnerAlias 镜像来源
type ImageOwnerAlias string

// 镜像来源
const (
	SystemImageOwnerAlias      ImageOwnerAlias = "system"
	SelfImageOwnerAlias        ImageOwnerAlias = "self"
	OthersImageOwnerAlias      ImageOwnerAlias = "others"
	MarketplaceImageOwnerAlias ImageOwnerAlias = "marketplace"
)

// Image 镜像
type Image struct {
	*params.CostInfo

	ID        string    `json:"_id"`
	RegionID  string    `json:"region_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	InstanceID   *string `json:"instance_id"`
	InstanceName *string `json:"instance_name"`
	RootImageID  *string `json:"root_image_id"`

	ImageInfo ImageInfo `json:"image_info"`
}

// SecurityGroup 安全组
type SecurityGroup struct {
	ID                        string                   `json:"_id"`
	UID                       uint32                   `json:"uid"`
	RegionID                  string                   `json:"region_id"`
	SecurityGroupName         string                   `json:"security_group_name"`
	Description               string                   `json:"description"`
	VpcID                     string                   `json:"vpc_id"`
	VpcName                   string                   `json:"vpc_name"`
	ClientToken               string                   `json:"client_token"`
	SecurityGroupID           string                   `json:"security_group_id"`
	IsDefault                 bool                     `json:"is_default"` // 是否为默认安全组，只在新建的时候插入该字段
	InnerAccessPolicy         InnerAccessPolicy        `json:"inner_access_policy"`
	ReferencingSecurityGroups ReferencingSecurityGroup `json:"referencing_security_groups"`
	UpdatedAt                 time.Time                `json:"updated_at"`
	CreatedAt                 time.Time                `json:"created_at"`
}

// NicType 网络类型
type NicType string

// NicType 网络类型
const (
	InternetNicType NicType = "internet"
	IntranetNicType NicType = "intranet"
)

// IPProtocol 协议
type IPProtocol string

// IPProtocol 协议
const (
	AllIPProtocol  IPProtocol = "all"
	TCPIPProtocol  IPProtocol = "tcp"
	UDPIPProtocol  IPProtocol = "udp"
	ICMPIPProtocol IPProtocol = "icmp"
	GREIPProtocol  IPProtocol = "gre"
)

// PermissionPolicy 授权策略
type PermissionPolicy string

// PermissionPolicy 授权策略
const (
	AcceptPermissionPolicy PermissionPolicy = "accept"
	DropPermissionPolicy   PermissionPolicy = "drop"
)

// InnerAccessPolicy 内部访问策略
type InnerAccessPolicy string

// InnerAccessPolicy 内部访问策略
const (
	AcceptInnerAccessPolicy InnerAccessPolicy = "Accept"
	DropInnerAccessPolicy   InnerAccessPolicy = "Drop"
)

// SecurityGroupPriority 安全组优先级
type SecurityGroupPriority int

// IsValid .
func (p SecurityGroupPriority) IsValid() bool {
	if p >= 1 && p <= 100 {
		return true
	}
	return false
}

// Direction 策略方向
type Direction string

// Direction 策略方向
const (
	IngressDirection Direction = "ingress"
	EgressDirection  Direction = "egress"
	AllDirection     Direction = "all"
)

// ReferencingSecurityGroup 正在授权给这个安全组的其他安全组信息列表
type ReferencingSecurityGroup struct {
	ReferencingSecurityGroup []ReferencingSecurityGroupType `json:"referencing_security_group"`
}

// ReferencingSecurityGroupType 正在授权给这个安全组的其他安全组信息
type ReferencingSecurityGroupType struct {
	SecurityGroupID string `json:"security_group_id"` // 安全组ID
	ProviderUID     int    `json:"ali_uid"`           // 这个安全组所属用户ID
}

// PermissionType 安全组权限规则
type PermissionType struct {
	SecurityGroupRuleID     string                `json:"security_group_rule_id"`
	IPProtocol              IPProtocol            `json:"ip_protocol"`
	PortRange               string                `json:"port_range"`
	SourceCidrIP            string                `json:"source_cidr_ip"`
	SourceGroupID           string                `json:"source_group_id"`
	SourceGroupName         string                `json:"source_group_name"`
	SourceGroupOwnerAccount string                `json:"source_group_owner_account"`
	DestCidrIP              string                `json:"dest_cidr_ip"`
	DestGroupID             string                `json:"dest_group_id"`
	DestGroupName           string                `json:"dest_group_name"`
	DestGroupOwnerAccount   string                `json:"dest_group_owner_account"`
	Policy                  PermissionPolicy      `json:"policy"`
	NicType                 NicType               `json:"nic_type"`
	Priority                SecurityGroupPriority `json:"priority"`
	Direction               Direction             `json:"direction"`
	Description             string                `json:"description"`
	CreateTime              string                `json:"create_time"`
	IPv6SourceCidrIP        string                `json:"ipv6_source_cidr_ip"`
	IPv6DestCidrIP          string                `json:"ipv6_dest_cidr_ip"`
	SourcePortRange         string                `json:"source_port_range"`
}

// AuthType 源类型
type AuthType string

// ip段和安全组id的源类型
const (
	AuthIPCidr          AuthType = "ip_cidr"
	AuthSecurityGroupID AuthType = "security_group_id"
)
