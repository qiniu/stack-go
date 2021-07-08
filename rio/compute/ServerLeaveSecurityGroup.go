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

// ServerLeaveSecurityGroup ..
func (d *Server) ServerLeaveSecurityGroup(args *ServerLeaveSecurityGroupArgs) (resp *ServerJoinSecurityGroupResp, err error) {
	str := "/api/rio/v1/compute/server"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/security_group/leave", args.ServerID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
