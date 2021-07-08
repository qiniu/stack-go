package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerModifyVNCPwdArgs 参数
type ServerModifyVNCPwdArgs struct {
	ZoneID      string `json:"zone-id"`
	ServerID    string `json:"server_id"`
	VNCPassword string `json:"vnc_password"`
}

// ServerModifyVNCPwdResp 返回
type ServerModifyVNCPwdResp struct {
	common.Response
}

// ServerModifyVNCPwd ..
func (d *Server) ServerModifyVNCPwd(args *ServerModifyVNCPwdArgs) (resp *ServerModifyVNCPwdResp, err error) {
	str := "/api/rio/v1/compute/server"
	req := client.NewRequest(http.MethodPut, fmt.Sprintf(str+"/%s/vnc/password", args.ServerID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
