package vpc

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// CreateEIPParams EIP创建参数
type CreateEIPParams struct {
	RegionID       string  `json:"region_id"`
	AllocationName string  `json:"allocation_name"` // 分配名称
	Bandwidth      *string `json:"bandwidth"`       // EIP的带宽峰值，取值范围：1~200，单位为Mbps。 默认值为5。
	ISP            *EIPISP `json:"isp"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`

	EnableAutoRenew       *bool   `json:"enable_auto_renew,omitempty"`       // 是否打开自动续费
	CustomizedReleaseTime *string `json:"customized_release_time,omitempty"` //后付费定时删除的时间
	*params.CostParams
}

// CreateEIPResponse EIP创建返回数据
type CreateEIPResponse struct {
	RequestID string `json:"request_id"`
	Data      *struct {
		AllocationID *string `json:"allocation_id"` // 按量创建返回 ID
		OrderHash    *string `json:"order_hash"`    // 预付费创建返回订单号
	} `json:"data"`
}

// CreateEIP EIP创建
func (s *VPC) CreateEIP(p *CreateEIPParams) (resp *CreateEIPResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/eip").WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
