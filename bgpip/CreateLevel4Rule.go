package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateLevel4RuleParams BGP高防四层规则创建参数
type CreateLevel4RuleParams struct {
	RegionID    string         `json:"region_id"`   // 地域 id
	ResourceID  string         `json:"resource_id"` // 高防实例 id
	RuleName    string         `json:"rule_name"`
	Protocol    Level4Protocol `json:"protocol"`
	VirtualPort uint16         `json:"virtual_port"` // 转发端口
	SourcePort  uint16         `json:"source_port"`  // 源站端口
	SourceType  SourceType     `json:"source_type"`  // 回源方式，取值[1(域名回源)，2(IP回源)]
	KeepTime    uint32         `json:"keep_time"`    // 会话保持时间，单位秒
	KeepEnable  SwitchStatus   `json:"keep_enable"`  // 会话保持开关，取值[0(会话保持关闭)，1(会话保持开启)]
	SourceList  []L4RuleSource `json:"source_list"`  // 回源列表
}

// Level4Protocol 四层协议
type Level4Protocol string

// TCP和UDP协议
const (
	Level4ProtocolTCP Level4Protocol = "TCP"
	Level4ProtocolUDP Level4Protocol = "UDP"
)

// SourceType 源站类型
type SourceType uint8

// 域名和IP源站
const (
	SourceTypeDOMAIN SourceType = 1
	SourceTypeIP     SourceType = 2
)

// SwitchStatus 开关状态
type SwitchStatus uint8

// 开关常量
const (
	SwitchStatusOn  SwitchStatus = 1
	SwitchStatusOff SwitchStatus = 0
)

// L4RuleSource 四层规则信息
type L4RuleSource struct {
	Source string `tson:"Source" json:"source"` // 回源地址
	Weight int    `tson:"Weight" json:"weight"` // 权重
}

// CreateLevel4RuleResponse BGP高防四层规则创建返回数据
type CreateLevel4RuleResponse struct {
	RequestID string `json:"request_id"`
}

// CreateLevel4Rule BGP高防四层规则创建
func (s *BGPIP) CreateLevel4Rule(p *CreateLevel4RuleParams) (resp *CreateLevel4RuleResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/security/bgpip/resource/%s/level4rule_v2/create", p.ResourceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
