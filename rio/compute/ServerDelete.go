package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerDeleteArgs 主机删除参数
type ServerDeleteArgs struct {
	ZoneID   string `json:"zone_id"`
	ServerID string `json:"server_id"`
	Force    *bool  `json:"force"`
}

// ServerDeleteResp 主机删除返回
type ServerDeleteResp struct {
	common.Response
}

// ServerDelete 主机删除
func (d *Server) ServerDelete(args *ServerDeleteArgs) (resp *SecurityGroupDeleteResp, err error) {
	str := "/api/rio/v1/compute/server"
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf(str+"/%s", args.ServerID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
