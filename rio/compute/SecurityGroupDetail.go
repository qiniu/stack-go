package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupDetailArgs 安全组详情参数
type SecurityGroupDetailArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupID string `json:"securitygroup_id"`
}

// SecurityGroupDetailResp 安全组详情返回
type SecurityGroupDetailResp struct {
	common.Response
	Data *SecurityGroup `json:"data"`
}

// SecurityGroupDetail 安全组详情
func (d *SecurityGroup) SecurityGroupDetail(args *SecurityGroupDetailArgs) (resp *SecurityGroupDetailResp, err error) {
	str := "/api/rio/v1/compute/security_group"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/%s/detail", args.SecurityGroupID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
