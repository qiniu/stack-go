package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerModifyVNCPwdArgs 主机VNC密码修改参数
type ServerModifyVNCPwdArgs struct {
	ZoneID      string `json:"zone_id"`
	ServerID    string `json:"server_id"`
	VNCPassword string `json:"vnc_password"`
}

// ServerModifyVNCPwdResp 主机VNC密码修改返回
type ServerModifyVNCPwdResp struct {
	common.Response
}

// ServerModifyVNCPwd 主机修改VNC密码
func (s *Server) ServerModifyVNCPwd(args *ServerModifyVNCPwdArgs) (resp *ServerModifyVNCPwdResp, err error) {
	url := fmt.Sprintf("%s/server/%s/vnc/password", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPut, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
