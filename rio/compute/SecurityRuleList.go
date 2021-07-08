package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupRuleListArgs 参数
type SecurityGroupRuleListArgs struct {
	Marker string `json:"marker"`
	Limit  int    `json:"limit"`
	ZoneID string `json:"zone_id"`

	SecurityGroupID string `pos:"path:id"`
}

// SecurityGroupRuleListResp 返回
type SecurityGroupRuleListResp struct {
	common.Response
	Data []*SecurityGroupRule `json:"data"`
}

// SecurityGroupRuleList .
func (d *SecurityGroupRule) SecurityGroupRuleList(args *SecurityGroupRuleListArgs) (resp *SecurityGroupRuleListResp, err error) {
	str := "/api/rio/v1/compute/security_group"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/%s/rules", args.SecurityGroupID)).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
