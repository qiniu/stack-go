package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupRuleListArgs 安全组规则列表参数
type SecurityGroupRuleListArgs struct {
	Marker          string `json:"marker"`
	Limit           int    `json:"limit"`
	ZoneID          string `json:"zone_id"`
	SecurityGroupID string `json:"security_group_id"`
}

// SecurityGroupRuleListResp 安全组规则列表返回
type SecurityGroupRuleListResp struct {
	common.Response
	Data []*SecurityGroupRule `json:"data"`
}

// SecurityGroupRuleList 安全组规则列表
func (s *SecurityGroupRule) SecurityGroupRuleList(args *SecurityGroupRuleListArgs) (resp *SecurityGroupRuleListResp, err error) {
	url := fmt.Sprintf("%s/security_group/%s/rules", ComputURLPrefix, args.SecurityGroupID)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
