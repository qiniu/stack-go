package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeLevel4RulesParams BGP高防四层规则列表参数
type DescribeLevel4RulesParams struct {
	RegionID   string `json:"region_id"`      // 地域 id
	ResourceID string `json:"resource_id"`    // 高防实例 id
	Page       *int   `json:"page,omitempty"` // 查询页码
	Size       *int   `json:"size,omitempty"` // 分页大小
}

// DescribeLevel4RulesResponse BGP高防四层规则列表返回数据
type DescribeLevel4RulesResponse struct {
	Page  int                `json:"page"`  // 查询页码
	Size  int                `json:"size"`  // 分页大小
	Total int                `json:"total"` // 总数
	Data  []Level4RuleDetail `json:"data"`  // 规则列表
}

// Level4RuleDetail 四层规则详情
type Level4RuleDetail struct {
	RuleName    string          `json:"rule_name"`    // 规则名称
	RuleID      string          `json:"rule_id"`      // 规则 id
	Protocol    string          `json:"protocol"`     // 四层协议，取值：TCP, UDP
	VirtualPort int             `json:"virtual_port"` // 转发端口
	SourcePort  int             `json:"source_port"`  // 源站端口
	SourceType  int             `json:"source_type"`  // 回源方式，取值[1(域名回源)，2(IP回源)]
	KeepTime    int             `json:"keep_time"`    // 会话保持时间，单位秒
	KeepEnable  int             `json:"keep_enable"`  // 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	SourceList  []L4RuleSource  `json:"source_list"`  // 回源列表
	HealthCheck HealthCheckInfo `json:"health_check"`
}

// HealthCheckInfo 四层健康检查信息
type HealthCheckInfo struct {
	Protocol    string `json:"protocol"`     //  四层协议，取值：TCP, UDP
	VirtualPort int    `json:"virtual_port"` // 转发端口
	RuleID      string `json:"rule_id"`      // 规则id
	Enable      int    `json:"enable"`       // 开启状态，1表示开启，0表示关闭
	Interval    int    `json:"interval"`     // 检测间隔时间，单位秒, 检查间隔必须大于响应超时
	KickNum     int    `json:"kick_num"`     // 不健康阈值，单位次
	AliveNum    int    `json:"alive_num"`    // 健康阈值，单位次
	TimeOut     int    `json:"timeout"`      // 响应超时时间，单位秒
	ErrMsg      string `json:"err_msg"`      // 健康检查错误信息
}

// DescribeLevel4Rules BGP高防四层规则列表
func (s *BGPIP) DescribeLevel4Rules(p *DescribeLevel4RulesParams) (resp *DescribeLevel4RulesResponse, err error) {
	req := client.NewRequest(http.MethodGet, fmt.Sprintf("/v1/security/bgpip/resource/%s/level4rule_v2", p.ResourceID)).WithRegionID(&p.RegionID).WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
