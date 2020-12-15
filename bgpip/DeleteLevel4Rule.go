package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteLevel4RuleParams BGP高防四层规则删除参数
type DeleteLevel4RuleParams struct {
	RegionID   string `json:"region_id"`   // 地域 id
	ResourceID string `json:"resource_id"` // 高防实例 id
	RuleID     string `json:"rule_id"`     // 四层规则 id
}

// DeleteLevel4RuleResponse BGP高防四层规则删除返回数据
type DeleteLevel4RuleResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteLevel4Rule BGP高防四层规则删除
func (s *BGPIP) DeleteLevel4Rule(p *DeleteLevel4RuleParams) (resp *DeleteLevel4RuleResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/security/bgpip/resource/%s/level4rule_v2/delete", p.ResourceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
