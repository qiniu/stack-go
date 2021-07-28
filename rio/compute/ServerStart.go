package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerStartArgs 开机参数
type ServerStartArgs struct {
	ZoneID   string `json:"zone_id"`
	ServerID string `json:"server_id"`
}

// ServerStartResp 开机返回
type ServerStartResp struct {
	common.Response
}

// ServerStart 开机
func (s *Server) ServerStart(args *ServerStartArgs) (resp *ServerStartResp, err error) {
	url := fmt.Sprintf("%s/server/%s/start", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
