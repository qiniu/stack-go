package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerVNCArgs 参数
type ServerVNCArgs struct {
	ZoneID   string `json:"zone_id"`
	ServerID string `json:"server_id"`
}

// ServerVNCResp 返回
type ServerVNCResp struct {
	common.Response
	URL string `json:"data"`
}

// ServerVNC ..
func (d *Server) ServerVNC(args *ServerVNCArgs) (resp *ServerVNCResp, err error) {
	url := fmt.Sprintf("%s/server/%s/vnc", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
