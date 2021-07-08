package compute

import (
	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// Server ..
type Server struct {
	client *client.Client
}

// NewServer 获得主机实例
func NewServer(cli *client.Client) *Server {
	return &Server{client: cli}
}

//ServerInfo 主机输入封装
type ServerInfo struct {
	ServerID       string                     `json:"server_id"`
	ServerName     string                     `json:"server_name"`
	Description    string                     `json:"description"`
	ZoneID         string                     `json:"zone_id"`
	Status         string                     `json:"status"`
	Hostname       string                     `json:"hostname"`
	Tags           []*TagInfo                 `json:"tags"`
	Image          *ServerImageInfo           `json:"image"`
	Flavor         *ServerFlavorInfo          `json:"flavor"`
	VPC            []*ServerVPCInfo           `json:"vpc"`
	Disks          []*ServerDiskInfo          `json:"disks"`
	SecurityGroups []*ServerSecurityGroupInfo `json:"security_groups"`
	KeyPairName    string                     `json:"key_pair_name"`
	ExpiredAt      int64                      `json:"expired_at"`
	CreatedAt      int64                      `json:"created_at"`
}

//ServerDiskInfo 封装
type ServerDiskInfo struct {
	ID     string `json:"id"`
	Size   int    `json:"size"`    // GB
	IsBoot bool   `json:"is_boot"` // 是否是系统盘
}

//ServerImageInfo 封装
type ServerImageInfo struct {
	ImageID   string `json:"image_id"`
	ImageName string `json:"image_name"`
	OSType    string `json:"os_type"`
	OSName    string `json:"os_name"`
}

//ServerFlavorInfo 封装
type ServerFlavorInfo struct {
	FlavorID   string `json:"flavor_id"`
	FlavorName string `json:"flavor_name"`
	CPU        int64  `json:"cpu"`
	Memory     int64  `json:"memory"`               // MB
	GPUSpec    string `json:"gpu_spec,omitempty"`   // 实例规格附带的GPU类型。
	GPUAmount  int64  `json:"gpu_amount,omitempty"` // 实例规格附带的GPU数量。
}

//ServerEIPInfo 封装
type ServerEIPInfo struct {
	EIPID     string `json:"eip_id"`
	EIPName   string `json:"eip_name"`
	IPAddress string `json:"ip_address"`
	Bandwidth int64  `json:"bandwidth"` // Mbps
}

//ServerVPCInfo 封装
type ServerVPCInfo struct {
	PrivateIPAddress string         `json:"private_ip_address"`
	VPCID            string         `json:"vpc_id"`
	VPCName          string         `json:"vpc_name"`
	VSwitchID        string         `json:"v_switch_id"`
	VSwitchName      string         `json:"v_switch_name"`
	EipInfo          *ServerEIPInfo `json:"eip_info,omitempty"`
}

//ServerSecurityGroupInfo 封装
type ServerSecurityGroupInfo struct {
	SecurityGroupID   string `json:"security_group_id"`
	SecurityGroupName string `json:"security_group_name"`
}

//ServerDiskArgs 封装
type ServerDiskArgs struct {
	Size             int                  `json:"size,omitempty"`
	Category         *common.DiskCategory `json:"category,omitempty"`
	SnapshotID       *string              `json:"snapshot_id,omitempty"`
	DiskName         *string              `json:"disk_name,omitempty"`
	Description      *string              `json:"description,omitempty"`
	DeleteWithServer *bool                `json:"delete_with_server,omitempty"`
}
