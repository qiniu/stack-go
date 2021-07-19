package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerLeaveSecurityGroupArgs 参数
type ServerLeaveSecurityGroupArgs struct {
	ZoneID          string `json:"zone_id"`
	ServerID        string `json:"server_id"`
	SecurityGroupID string `json:"security_group_id"`
}

// ServerLeaveSecurityGroupResp 返回
type ServerLeaveSecurityGroupResp struct {
	common.Response
}

// ServerLeaveSecurityGroup 主机分离安全组
func (s *Server) ServerLeaveSecurityGroup(args *ServerLeaveSecurityGroupArgs) (resp *ServerJoinSecurityGroupResp, err error) {
	url := fmt.Sprintf("%s/server/%s/security_group/leave", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
