package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// AssociateEIPParams 弹性公网IP绑定参数
type AssociateEIPParams struct {
	// 要绑定云产品实例的 EIP 所属的地域ID。
	RegionID string `json:"region_id"`

	// 绑定云产品实例的EIP的ID。
	AllocationID string `json:"allocation_id"`

	// 要绑定EIP的实例的类型，取值：
	// - Nat：NAT网关。
	// - SlbInstance：负载均衡SLB。
	// - EcsInstance：云服务器ECS。
	// - NetworkInterface：辅助弹性网卡。
	// - HaVip：高可用虚拟IP。
	InstanceType *InstanceType `json:"instance_type"`

	// 要绑定EIP的实例ID。
	//
	// 支持输入NAT网关实例ID、负载均衡SLB实例ID、云服务器ECS实例ID、辅助弹性网卡实例ID、高可用虚拟IP实例ID。
	InstanceID string `json:"instance_id"`
}

// AssociateEIPResponse 弹性公网IP绑定返回数据
type AssociateEIPResponse struct {
	RequestID string `json:"request_id"`
}

// AssociateEIP 弹性公网IP绑定
func (s *VPC) AssociateEIP(p *AssociateEIPParams) (resp *AssociateEIPResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/eip/%s/associate", p.AllocationID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
