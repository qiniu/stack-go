package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupRuleDeleteArgs ..
type SecurityGroupRuleDeleteArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupID     string `json:"security_group_id"`
	SecurityGroupRuleID string `json:"security_group_rule_id"`
}

// SecurityGroupRuleDeleteResp ..
type SecurityGroupRuleDeleteResp struct {
	common.Response
}

// SecurityGroupRuleDelete ..
func (d *SecurityGroupRule) SecurityGroupRuleDelete(args *SecurityGroupRuleDeleteArgs) (resp *SecurityGroupRuleDeleteResp, err error) {
	url := fmt.Sprintf("%s/security_group/rule/%s", ComputURLPrefix, args.SecurityGroupRuleID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
