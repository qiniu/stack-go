package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyLevel4RuleParams BGP高防四层规则修改参数
type ModifyLevel4RuleParams struct {
	RegionID    string         `json:"region_id"`    // 地域 id
	ResourceID  string         `json:"resource_id"`  // 高防实例 id
	RuleID      string         `json:"rule_id"`      // 四层规则 id
	Protocol    Level4Protocol `json:"protocol"`     // 四层协议，取值：TCP, UDP
	RuleName    string         `json:"rule_name"`    // 规则名称
	VirtualPort uint16         `json:"virtual_port"` // 转发端口
	SourcePort  uint16         `json:"source_port"`  // 源站端口
	SourceType  SourceType     `json:"source_type"`  // 回源方式，取值[1(域名回源)，2(IP回源)]
	SourceList  []L4RuleSource `json:"source_list"`  // 回源列表
	LbType      int            `json:"lb_type"`      // 负载均衡方式，取值: 1(加权轮询)
}

// ModifyLevel4RuleResponse BGP高防四层规则修改返回数据
type ModifyLevel4RuleResponse struct {
	RequestID string `json:"request_id"`
}

// ModifyLevel4Rule BGP高防四层规则修改
func (s *BGPIP) ModifyLevel4Rule(p *ModifyLevel4RuleParams) (resp *ModifyLevel4RuleResponse, err error) {
	req := client.NewRequest(http.MethodPut, fmt.Sprintf("/v1/security/bgpip/resource/%s/level4rule_v2/%s/update", p.ResourceID, p.RuleID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
