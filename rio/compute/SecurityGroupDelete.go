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
func (d *SecurityGroup) SecurityGroupDelete(args *SecurityGroupDeleteArgs) (resp *SecurityGroupDeleteResp, err error) {
	str := "/api/rio/v1/compute/security_group"
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf(str+"/%s", args.SecurityGroupID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
