package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupDeleteArgs 删除安全组参数
type SecurityGroupDeleteArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupID string `json:"securitygroup_id"`
}

// SecurityGroupDeleteResp 删除安全组返回
type SecurityGroupDeleteResp struct {
	common.Response
}

// SecurityGroupDelete 删除安全组
func (s *SecurityGroup) SecurityGroupDelete(args *SecurityGroupDeleteArgs) (resp *SecurityGroupDeleteResp, err error) {
	url := fmt.Sprintf("%s/security_group/%s", ComputURLPrefix, args.SecurityGroupID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
