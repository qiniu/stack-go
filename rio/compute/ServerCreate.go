package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerCreateArgs 主机创建参数
type ServerCreateArgs struct {
	ZoneID          string           `json:"zone_id"`
	ImageID         string           `json:"image_id,omitempty"`
	FlavorID        string           `json:"flavor_id,omitempty"`
	SecurityGroupID string           `json:"security_group_id,omitempty"`
	VSwitchID       *string          `json:"v_switch_id,omitempty"`
	ServerName      string           `json:"server_name,omitempty"`
	HostName        *string          `json:"host_name,omitempty"`
	Description     *string          `json:"description,omitempty"`
	Password        string           `json:"password,omitempty"`
	UserData        *string          `json:"user_data,omitempty"`
	KeyPairName     string           `json:"key_pair_name,omitempty"`
	VPCID           *string          `json:"vpcid,omitempty"`
	SystemDisk      ServerDiskArgs   `json:"system_disk,omitempty"`
	DataDisks       []ServerDiskArgs `json:"data_disks,omitempty"`
	ClientToken     *string          `json:"client_token,omitempty"`
}

// ServerCreateResp 主机创建返回
type ServerCreateResp struct {
	common.Response
}

// ServerCreate 主机创建
func (s *Server) ServerCreate(args *ServerCreateArgs) (resp *ServerCreateResp, err error) {
	url := fmt.Sprintf("%s/server", ComputURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
