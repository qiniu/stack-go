package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerJoinSecurityGroupArgs 参数
type ServerJoinSecurityGroupArgs struct {
	ZoneID          string `json:"zone_id"`
	ServerID        string `json:"server_id"`
	SecurityGroupID string `json:"security_group_id"`
}

// ServerJoinSecurityGroupResp 返回
type ServerJoinSecurityGroupResp struct {
	common.Response
}

// ServerJoinSecurityGroup 主机组合安全组
func (d *Server) ServerJoinSecurityGroup(args *ServerJoinSecurityGroupArgs) (resp *ServerJoinSecurityGroupResp, err error) {
	url := fmt.Sprintf("%s/server/%s", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
