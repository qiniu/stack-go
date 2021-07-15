package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerModifyAttrArgs 参数
type ServerModifyAttrArgs struct {
	ZoneID      string  `json:"zone_id"`
	ServerID    string  `json:"server_id"`
	ServerName  string  `json:"server_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Password    string  `json:"password,omitempty"`
	HostName    *string `json:"host_name,omitempty"`
	UserData    *string `json:"user_data,omitempty"`
}

// ServerModifyAttrResp 返回
type ServerModifyAttrResp struct {
	common.Response
}

// ServerModifyAttr 主机修改
func (d *Server) ServerModifyAttr(args *ServerModifyAttrArgs) (resp *ServerModifyAttrResp, err error) {
	url := fmt.Sprintf("%s/server/%s", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPatch, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
