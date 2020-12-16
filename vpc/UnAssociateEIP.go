package vpc

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// UnAssociateEIPParams 弹性公网IP解绑参数
type UnAssociateEIPParams struct {
	// 要解绑的EIP的地域ID。
	RegionID string `json:"region_id"`

	// 要解绑的EIP的ID。
	AllocationID string `json:"allocation_id"`

	// 要解绑EIP的云产品类型，取值：
	// - EcsInstance（默认值）：专有网络类型的ECS实例。
	// - SlbInstance：专有网络类型的SLB实例。
	// - NetworkInterface：专有网络类型的辅助弹性网卡。
	// - Nat：NAT网关。
	// - HaVip：高可用虚拟IP。
	InstanceType *InstanceType `json:"instance_type"`

	// 要解绑EIP的云产品实例的ID。
	InstanceID string `json:"instance_id"`

	// 当EIP绑定了NAT网关，且NAT网关添加了DNAT或SNAT条目时，是否强制解绑EIP，取值：
	// - false（默认值）：不强制解绑EIP。
	// - true：强制解绑EIP。
	Force *bool `json:"force,omitempty"`
}

// UnAssociateEIPResponse 弹性公网IP解绑返回数据
type UnAssociateEIPResponse struct {
	RequestID string `json:"request_id"`
}

// UnAssociateEIP 弹性公网IP解绑
func (s *VPC) UnAssociateEIP(p *UnAssociateEIPParams) (resp *UnAssociateEIPResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/eip/%s/associate", p.AllocationID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
