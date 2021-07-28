package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerModifyAttrArgs 主机属性修改参数
type ServerModifyAttrArgs struct {
	ZoneID      string  `json:"zone_id"`
	ServerID    string  `json:"server_id"`
	ServerName  string  `json:"server_name,omitempty"`
	Description *string `json:"description,omitempty"`
	Password    string  `json:"password,omitempty"`
	HostName    *string `json:"host_name,omitempty"`
	UserData    *string `json:"user_data,omitempty"` // 实例自定义数据
}

// ServerModifyAttrResp 主机属性修改返回
type ServerModifyAttrResp struct {
	common.Response
}

// ServerModifyAttr 主机属性修改
func (s *Server) ServerModifyAttr(args *ServerModifyAttrArgs) (resp *ServerModifyAttrResp, err error) {
	url := fmt.Sprintf("%s/server/%s", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPatch, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
