package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupModifyArgs 安全组修改参数
type SecurityGroupModifyArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupID   string  `json:"security_group_id"`
	SecurityGroupName string  `json:"security_group_name"`
	Description       *string `json:"description"`
}

// SecurityGroupModifyResp 安全组返回参数
type SecurityGroupModifyResp struct {
	common.Response
}

// SecurityGroupModify 安全组修改
func (d *SecurityGroup) SecurityGroupModify(args *SecurityGroupModifyArgs) (resp *SecurityGroupModifyResp, err error) {
	url := fmt.Sprintf("%s/security_group/%s", ComputURLPrefix, args.SecurityGroupID)
	req := client.NewRequest(http.MethodPut, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
