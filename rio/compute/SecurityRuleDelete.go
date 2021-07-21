package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupRuleDeleteArgs 安全组规则删除参数
type SecurityGroupRuleDeleteArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupID     string `json:"security_group_id"`
	SecurityGroupRuleID string `json:"security_group_rule_id"`
}

// SecurityGroupRuleDeleteResp 返回
type SecurityGroupRuleDeleteResp struct {
	common.Response
}

// SecurityGroupRuleDelete 安全组规则删除
func (s *SecurityGroupRule) SecurityGroupRuleDelete(args *SecurityGroupRuleDeleteArgs) (resp *SecurityGroupRuleDeleteResp, err error) {
	url := fmt.Sprintf("%s/security_group/rule/%s", ComputURLPrefix, args.SecurityGroupRuleID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
