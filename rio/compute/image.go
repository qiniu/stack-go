package compute

import "github.com/qiniu/stack-go/components/client"

// Image 镜像类接口封装
type Image struct {
	client *client.Client
}

// NewImage 获得镜像实例
func NewImage(cli *client.Client) *Image {
	return &Image{client: cli}
}

// ImageContainer 容器格式
type ImageContainer string

// 镜像容器格式枚举
const (
	ImageContainerAMI    ImageContainer = "ami"
	ImageContainerARI    ImageContainer = "ari"
	ImageContainerAKI    ImageContainer = "aki"
	ImageContainerBARE   ImageContainer = "bare"
	ImageContainerOVF    ImageContainer = "ovf"
	ImageContainerOVA    ImageContainer = "ova"
	ImageContainerDOCKER ImageContainer = "docker"
)

// ImageDisk 镜像格式
type ImageDisk string

// 镜像格式枚举
const (
	ImageDiskAMI   ImageDisk = "ami"
	ImageDiskARI   ImageDisk = "ari"
	ImageDiskAKI   ImageDisk = "aki"
	ImageDiskVHD   ImageDisk = "vhd"
	ImageDiskVHDX  ImageDisk = "vhdx"
	ImageDiskVMDK  ImageDisk = "vmdk"
	ImageDiskRAW   ImageDisk = "raw"
	ImageDiskQCOW2 ImageDisk = "qcow2"
	ImageDiskVDI   ImageDisk = "vdi"
	ImageDiskPLOOP ImageDisk = "ploop"
	ImageDiskISO   ImageDisk = "iso"
)

// ImageVisibility 可见性类型
type ImageVisibility string

// 可见性枚举
const (
	ImageVisibilityPublic  ImageVisibility = "public"
	ImageVisibilityPrivate ImageVisibility = "private"
	ImageVisibilityShared  ImageVisibility = "shared"
)

// ImageStatus 镜像状态
type ImageStatus string

// 镜像状态枚举
const (
	// 排队状态，自定义创建镜像
	ImageStatusQueued ImageStatus = "queued"
	// 保存状态
	ImageStatusSaving ImageStatus = "saving"
	// 活动状态，正常可用镜像即为活动
	ImageStatusActive ImageStatus = "active"
	// 镜像上传发生错误
	ImageStatusKilled ImageStatus = "killed"
	// 已删除
	ImageStatusDeleted ImageStatus = "deleted"
	// 删除，未完成
	ImageStatusPendingDelete ImageStatus = "pending_delete"
	// 停用状态
	ImageStatusDeactivated ImageStatus = "deactivated"
	// 上传状态
	ImageStatusUploading ImageStatus = "uploading"
	// 导入状态
	ImageStatusImporting ImageStatus = "importing"
)

// PlatformType 操作系统平台类型
type PlatformType string

// 操作系统平台类型 Enum
const (
	PlatformTypeWindows  PlatformType = "Windows"
	PlatformTypeCentOS   PlatformType = "CentOS"
	PlatformTypeRedhat   PlatformType = "Redhat"
	PlatformTypeUbuntu   PlatformType = "Ubuntu"
	PlatformTypeSUSE     PlatformType = "SUSE"
	PlatformTypeDebian   PlatformType = "Debian"
	PlatformTypeFedora   PlatformType = "Fedora"
	PlatformTypeNeoshine PlatformType = "Neoshine Linux Server"
	PlatformTypeOracle   PlatformType = "Oracle Linux"
	PlatformTypeCustom   PlatformType = "凝思Linux"
	PlatformTypeAsianux  PlatformType = "红旗Linux"
	PlatformTypeKylin    PlatformType = "Kylin Linux"
	PlatformTypeUOS      PlatformType = "UOS"
	PlatformTypeCirrOS   PlatformType = "CirrOS"
	PlatformTypeOther    PlatformType = "Other Linux"
)

// ArchitectureType 镜像系统架构类型
type ArchitectureType string

// 镜像系统架构类型 Enum
const (
	ArchitectureType86    ArchitectureType = "x86_64"
	ArchitectureTypeARM   ArchitectureType = "arm"
	ArchitectureTypePPC64 ArchitectureType = "pcc64"
	ArchitectureTypeI386  ArchitectureType = "i386"
)
