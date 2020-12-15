package bgpip

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ModifyElasticLimitParams BGP高防弹性防护带宽修改参数
type ModifyElasticLimitParams struct {
	RegionID     string `json:"region_id"`   // 地域 id
	ResourceID   string `json:"resource_id"` // 高防实例 id
	ElasticLimit int    `json:"elstic_limit"`
	/***
	弹性防护带宽， 单位 Mbps, 取值[0 30000 40000 50000 60000 70000 80000 90000 100000 150000 200000 250000 300000],
	当取值为0时表示关闭弹性防护，否则取值需高于实例的保底防护带宽
	***/
}

// ModifyElasticLimitResponse BGP高防弹性防护带宽修改返回数据
type ModifyElasticLimitResponse struct {
	RequestID string `json:"request_id"`
}

// ModifyElasticLimit BGP高防弹性防护带宽修改
func (s *BGPIP) ModifyElasticLimit(p *ModifyElasticLimitParams) (resp *ModifyElasticLimitResponse, err error) {
	req := client.NewRequest(http.MethodPatch, fmt.Sprintf("/v1/security/bgpip/resource/%s/elastic_limit", p.ResourceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
