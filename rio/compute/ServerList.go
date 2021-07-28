package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerListArgs 主机列表参数
type ServerListArgs struct {
	Marker string `json:"marker"`
	Limit  int    `json:"limit"`

	ZoneID             string   `json:"zone_id"`
	Status             *string  `json:"status"`
	ServerName         *string  `json:"server_name"`
	VPCID              *string  `json:"vpc_id"`
	VSwitchID          *string  `json:"v_switch_id"`
	ImageID            *string  `json:"image_id"`
	FlavorID           *string  `json:"flavor_id"`
	ServerIDs          []string `json:"server_ids"`
	SecurityGroupID    *string  `json:"security_group_id"`
	PrivateIPAddresses []string `json:"private_ip_addresses"`
	PublicIPAddresses  []string `json:"public_ip_addresses"`
}

// ServerListResp 主机列表返回
type ServerListResp struct {
	common.Response
	Data []*ServerInfo `json:"data"`
}

// ServerList 主机列表
func (s *Server) ServerList(args *ServerListArgs) (resp *ServerDetailResp, err error) {
	url := fmt.Sprintf("%s/server", ComputURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
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
	Flavor         *ServerFlavorInfo          `json:"flavor"` // 主机规格信息
	VPC            []*ServerVPCInfo           `json:"vpc"`
	Disks          []*ServerDiskInfo          `json:"disks"`
	SecurityGroups []*ServerSecurityGroupInfo `json:"security_groups"`
	KeyPairName    string                     `json:"key_pair_name"`
	ExpiredAt      int64                      `json:"expired_at"`
	CreatedAt      int64                      `json:"created_at"`
}

// ServerDiskInfo 主机的磁盘信息
type ServerDiskInfo struct {
	ID     string `json:"id"`
	Size   int    `json:"size"`    // GB
	IsBoot bool   `json:"is_boot"` // 是否是系统盘
}

// ServerImageInfo 主机的镜像信息
type ServerImageInfo struct {
	ImageID   string `json:"image_id"`
	ImageName string `json:"image_name"`
	OSType    string `json:"os_type"`
	OSName    string `json:"os_name"`
}

// ServerFlavorInfo 主机的规格信息
type ServerFlavorInfo struct {
	FlavorID   string `json:"flavor_id"`   // 实例规格
	FlavorName string `json:"flavor_name"` // 实例规格名称
	CPU        int64  `json:"cpu"`
	Memory     int64  `json:"memory"`               // MB
	GPUSpec    string `json:"gpu_spec,omitempty"`   // 实例规格附带的GPU类型。
	GPUAmount  int64  `json:"gpu_amount,omitempty"` // 实例规格附带的GPU数量。
}

// ServerEIPInfo 主机绑定的ip信息
type ServerEIPInfo struct {
	EIPID     string `json:"eip_id"`
	EIPName   string `json:"eip_name"`
	IPAddress string `json:"ip_address"`
	Bandwidth int64  `json:"bandwidth"` // Mbps
}

// ServerVPCInfo 主机的vpc信息
type ServerVPCInfo struct {
	PrivateIPAddress string         `json:"private_ip_address"`
	VPCID            string         `json:"vpc_id"`
	VPCName          string         `json:"vpc_name"`
	VSwitchID        string         `json:"v_switch_id"`
	VSwitchName      string         `json:"v_switch_name"`
	EipInfo          *ServerEIPInfo `json:"eip_info,omitempty"`
}

// ServerSecurityGroupInfo 主机的安全组信息
type ServerSecurityGroupInfo struct {
	SecurityGroupID   string `json:"security_group_id"`
	SecurityGroupName string `json:"security_group_name"`
}

// ServerDiskArgs 创建主机的磁盘参数
type ServerDiskArgs struct {
	Size             int                  `json:"size,omitempty"`
	Category         *common.DiskCategory `json:"category,omitempty"`
	SnapshotID       *string              `json:"snapshot_id,omitempty"`
	DiskName         *string              `json:"disk_name,omitempty"`
	Description      *string              `json:"description,omitempty"`
	DeleteWithServer *bool                `json:"delete_with_server,omitempty"`
}
