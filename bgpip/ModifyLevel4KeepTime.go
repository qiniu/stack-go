package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyLevel4KeepTimeParams BGP高防四层会话保持修改参数
type ModifyLevel4KeepTimeParams struct {
	RegionID   string       `json:"region_id"`   // 地域 id
	ResourceID string       `json:"resource_id"` // 高防实例 id
	RuleID     string       `json:"rule_id"`     // 四层规则 id
	KeepTime   uint32       `json:"keep_time"`   // 会话保持时间，单位:秒
	KeepEnable SwitchStatus `json:"keep_enable"` // 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
}

// ModifyLevel4KeepTimeResponse BGP高防四层会话保持修改返回数据
type ModifyLevel4KeepTimeResponse struct {
	RequestID string `json:"request_id"`
}

// ModifyLevel4KeepTime BGP高防四层会话保持修改
func (s *BGPIP) ModifyLevel4KeepTime(p *ModifyLevel4KeepTimeParams) (resp *ModifyLevel4KeepTimeResponse, err error) {
	req := client.NewRequest(http.MethodPut, fmt.Sprintf("/v1/security/bgpip/resource/%s/level4rule_v2/%s/keep_time", p.ResourceID, p.RuleID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
