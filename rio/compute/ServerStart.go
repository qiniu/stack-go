package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerStartArgs 开启主机参数
type ServerStartArgs struct {
	ZoneID   string `json:"zone_id"`
	ServerID string `json:"server_id"`
}

// ServerStartResp 开启主机返回
type ServerStartResp struct {
	common.Response
}

// ServerStart 开启主机
func (d *Server) ServerStart(args *ServerStartArgs) (resp *ServerStartResp, err error) {
	str := "/api/rio/v1/compute/server"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/start", args.ServerID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
