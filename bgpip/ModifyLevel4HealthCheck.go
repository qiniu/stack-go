package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyLevel4HealthCheckParams BGP高防四层健康检查修改参数
type ModifyLevel4HealthCheckParams struct {
	RegionID   string       `json:"region_id"`   // 地域 id
	ResourceID string       `json:"resource_id"` // 高防实例 id
	RuleID     string       `json:"rule_id"`     // 四层规则 id
	Enable     SwitchStatus `json:"enable"`      // 开启状态，1表示开启，0表示关闭
	Interval   uint32       `json:"interval"`    // 检测间隔时间，单位:秒, 检查间隔必须大于响应超时
	KickNum    uint32       `json:"kick_num"`    // 不健康阈值，单位:次
	AliveNum   uint32       `json:"alive_num"`   // 健康阈值，单位:次
	TimeOut    uint32       `json:"timeout"`     // 响应超时时间，单位:秒
}

// ModifyLevel4HealthCheckResponse BGP高防四层健康检查修改返回数据
type ModifyLevel4HealthCheckResponse struct {
	RequestID string `json:"request_id"`
}

// ModifyLevel4HealthCheck BGP高防四层健康检查修改
func (s *BGPIP) ModifyLevel4HealthCheck(p *ModifyLevel4HealthCheckParams) (resp *ModifyLevel4HealthCheckResponse, err error) {
	req := client.NewRequest(http.MethodPut, fmt.Sprintf("/v1/security/bgpip/resource/%s/level4rule_v2/%s/health_check", p.ResourceID, p.RuleID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
