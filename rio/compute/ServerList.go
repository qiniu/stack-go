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
