package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerDetailArgs 主机具体参数
type ServerDetailArgs struct {
	ZoneID   string `json:"zone_id"`
	ServerID string `json:"server_id"`
}

// ServerDetailResp 主机具体参数返回
type ServerDetailResp struct {
	common.Response
	Data *ServerInfo `json:"data"`
}

// ServerDetail 主机具体参数
func (d *Server) ServerDetail(args *ServerDetailArgs) (resp *ServerDetailResp, err error) {
	str := "/api/rio/v1/compute/server"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/%s", args.ServerID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
