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
func (s *Server) ServerDelete(args *ServerDeleteArgs) (resp *SecurityGroupDeleteResp, err error) {
	url := fmt.Sprintf("%s/server/%s", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
