package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerListArgs 主机列表参数
type ServerListArgs struct {
	Marker string `json:"query:marker"`
	Limit  int    `json:"query:limit"`

	ZoneID             string   `json:"header:x-rio-zone-id"`
	Status             *string  `json:"query:status"`
	ServerName         *string  `json:"query:server_name"`
	VPCID              *string  `json:"query:vpc_id"`
	VSwitchID          *string  `json:"query:v_switch_id"`
	ImageID            *string  `json:"query:image_id"`
	FlavorID           *string  `json:"query:flavor_id"`
	ServerIDs          []string `json:"query:server_ids"`
	SecurityGroupID    *string  `json:"query:security_group_id"`
	PrivateIPAddresses []string `json:"query:private_ip_addresses"`
	PublicIPAddresses  []string `json:"query:public_ip_addresses"`
}

// ServerListResp 主机列表返回
type ServerListResp struct {
	common.Response
	Data []*ServerInfo `json:"data"`
}

// ServerList 主机列表
func (d *Server) ServerList(args *ServerListArgs) (resp *ServerDetailResp, err error) {
	url := fmt.Sprintf("%s/server", ComputURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
