package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerRebootArgs 参数
type ServerRebootArgs struct {
	ZoneID    string `json:"zone_id"`
	ServerID  string `json:"server_id"`
	ForceStop *bool  `json:"force_stop"`
}

// ServerRebootResp 返回
type ServerRebootResp struct {
	common.Response
}

// ServerReboot ..
func (d *Server) ServerReboot(args *ServerRebootArgs) (resp *ServerRebootResp, err error) {
	url := fmt.Sprintf("%s/server/%s/reboot", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
