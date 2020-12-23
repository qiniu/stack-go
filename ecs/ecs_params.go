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

	ImageID      string `json:"image_id"`     // 镜像 ID
	ImageName    string `json:"image_name"`   // 镜像的名称
	Description  string `json:"description"`  // 镜像描述
	OSType       string `json:"os_type"`      // 操作系统类型。可能值：windows, linux
	OSName       string `json:"os_name"`      // 操作系统的中文显示名称。
	Architecture string `json:"architecture"` // 镜像的体系架构。取值范围：`i386`, `x86_64`

	// 镜像来源。取值范围：
	// - `system` ：七牛云提供的公共镜像。
	// - `self` ：您创建的自定义镜像。
	// - `others` ：其他七牛云用户共享给您的镜像。
	// - `marketplace` ：镜像市场提供的镜像。您查询到的云市场镜像可以直接使用，无需提前订阅。您需要自行留意云市场镜像的收费详情。
	ImageOwnerAlias string `json:"image_owner_alias"`

	Progress string `json:"progress"` //

	// 镜像是否已经运行在ECS实例中。取值范围：
	//
	// - instance：镜像处于运行状态，有ECS实例使用。
	// - none：镜像处于闲置状态，暂无ECS实例使用。
	Usage string `json:"usage"`

	// 镜像版本
	ImageVersion string `json:"image_version"`

	// 镜像的状态。可能值：
	//
	// - UnAvailable：不可用
	// - Available：可用
	// - Creating：创建中
	// - CreateFailed：创建失败
	Status ImageStatus `json:"status"`

	IsSelfShared         string `json:"is_self_shared"`          // 是否共享过该自定义镜像给其他用户
	Platform             string `json:"platform"`                // 操作系统平台
	Size                 int64  `json:"size"`                    // 镜像大小，单位GiB
	IsSupportCloudinit   bool   `json:"is_support_cloudinit"`    // 是否支持Cloud Init
	IsSupportIoOptimized bool   `json:"is_support_io_optimized"` // 是否可以在I/O优化实例上运行
	IsCopied             bool   `json:"is_copied"`               // 是否是拷贝的镜像
	IsSubscribed         bool   `json:"is_subscribed"`           // 是否订阅了该镜像的商品码对应的镜像商品的服务条款
	SnapshotID           string `json:"snapshot_id"`             // 根据某一快照ID创建的自定义镜像
	MinDiskSize          string `json:"min_disk_size"`           // 镜像要求的最小系统盘盘容量
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
	CreatingInstanceStatus  InstanceStatus = "Creating"  // 创建中
	ChangingInstanceStatus  InstanceStatus = "Changing"  // 变配中
	RunningInstanceStatus   InstanceStatus = "Running"   // 运行中
	StartingInstanceStatus  InstanceStatus = "Starting"  // 开机中
	StoppingInstanceStatus  InstanceStatus = "Stopping"  // 关机中
	StoppedInstanceStatus   InstanceStatus = "Stopped"   // 已关机
	RebootingInstanceStatus InstanceStatus = "Rebooting" // 重启中
	PendingInstanceStatus   InstanceStatus = "Pending"   // 挂起中
	DeletedInstanceStatus   InstanceStatus = "Deleted"   // 已删除
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
	ZoneID   string         `json:"zone_id"`   // 可用区ID。
	RegionID string         `json:"region_id"` // 区域 ID
	DiskID   string         `json:"disk_id"`   // 磁盘 ID
	DiskName string         `json:"disk_name"` // 磁盘名称
	Type     DiskType       `json:"type"`      // 磁盘类型
	Size     int64          `json:"size"`      // 云盘或本地盘大小，单位GiB。
	Status   DiskStatusType `json:"status"`    // 云盘状态。可能值：In_use, Available, Attaching, Detaching, Creating, ReIniting
	Device   string         `json:"device"`    // 云盘或本地盘挂载的实例的设备名，例如 `/dev/xvdb`。该参数仅在 Status 参数值为 `In_use` 时才有值，其他状态时为空。
	Category DiskCategory   `json:"category"`  // 云盘种类

	// 云盘或本地盘是否支持卸载。取值范围：
	// - true：支持。可以独立存在，且可以在可用区内自由挂载和卸载。
	// - false：不支持。不可以独立存在，且不可以在可用区内自由挂载和卸载。
	//
	// 以下类型块存储的Portable属性都为false，生命周期与实例等同：
	// - 本地盘
	// - 本地SSD盘
	// - 包年包月数据盘
	Portable     bool      `json:"portable"`
	Encrypted    bool      `json:"encrypted"`     // 是否只筛选出加密云盘。默认值：false
	Description  string    `json:"description"`   // 云盘或本地盘描述。
	ExpiredTime  time.Time `json:"expired_time"`  // 过期时间
	AttachedTime time.Time `json:"attached_time"` // 挂载时间
	DetachedTime time.Time `json:"detached_time"` // 卸载时间。

	// 云盘或本地盘的计费方式。
	//
	// 可能值：
	// - PrePaid：包年包月
	// - PostPaid：按量付费
	DiskChargeType string `json:"disk_charge_type"`

	EnableAutoSnapshot bool `json:"enable_auto_snapshot"` // 云盘是否启用自动快照策略功能。

	// 是否同时删除自动快照。
	//
	// 可能值：
	// - true：删除云盘上的快照。
	// - false：保留云盘上的快照。
	DeleteAutoSnapshot bool `json:"delete_auto_snapshot"`

	// 云盘是否设置了自动快照策略。
	EnableAutomatedSnapshotPolicy bool `json:"enable_automated_snapshot_policy"`

	// 是否随实例释放。可能值：
	//
	// - true：释放实例时，这块云盘随实例一起释放。
	// - false：释放实例时，这块云盘保留不释放。
	DeleteWithInstance bool `json:"delete_with_instance"`

	// 创建ECS实例时使用的镜像ID，只有通过镜像创建的云盘才有值，否则为空。这个值在云盘的生命周期内始终不变。
	ImageID string `json:"image_id"`

	// 云盘或本地盘挂载的实例ID。
	//
	// 该参数值仅在Status参数值为In_use时才有值，其他状态时为空。
	InstanceID string `json:"instance_id"`

	// 云盘或本地盘挂载的实例名称
	InstanceName string `json:"instance_name"`

	// 创建云盘使用的快照ID。
	//
	// 如果创建云盘时，没有指定快照，则该参数值为空。该参数值在云盘的生命周期内始终不变。
	SourceSnapshotID string `json:"source_snapshot_id"`

	// 云盘采用的自动快照策略ID。
	AutoSnapshotPolicyID string `json:"auto_snapshot_policy_id"`

	// ESSD云盘的性能等级。
	//
	// 可能值：PL1, PL2, PL3
	PerformanceLevel ESSDPerformanceLevel `json:"performance_level"`

	// 每秒读写（I/O）操作的次数，单位：次/s。
	IOPS int64 `json:"iops"`

	// 每秒读操作的次数，单位：次/s。
	IOPSRead int64 `json:"iops_read"`

	// 每秒写操作的次数，单位：次/s。
	IOPSWrite int64 `json:"iops_write"`

	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间

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
	*params.CostInfo // 付费信息

	RegionID  string    `json:"region_id"`  // 实例所属的地域 ID。
	CreatedAt time.Time `json:"created_at"` // 创建时间
	UpdatedAt time.Time `json:"updated_at"` // 更新时间

	InstanceID   *string `json:"instance_id"`   // 实例 ID
	InstanceName *string `json:"instance_name"` // 实例名称

	ImageInfo // 镜像信息
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
