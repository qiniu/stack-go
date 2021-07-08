package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerStopArgs 关闭主机参数
type ServerStopArgs struct {
	ZoneID    string `json:"zone_id"`
	ServerID  string `json:"server_id"`
	ForceStop *bool  `json:"force_stop"`
}

// ServerStopResp 关闭主机返回
type ServerStopResp struct {
	common.Response
}

// ServerStop 关闭主机
func (d *Server) ServerStop(args *ServerStopArgs) (resp *ServerStopResp, err error) {
	str := "/api/rio/v1/compute/server"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/stop", args.ServerID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
