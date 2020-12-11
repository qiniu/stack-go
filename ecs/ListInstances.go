package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/params"
)

// ListInstancesParams 查询主机列表参数
type ListInstancesParams struct {
	Page             *int               `json:"page"`
	Size             *int               `json:"size"`
	InstanceID       *string            `json:"instance_id"`
	InstanceName     *string            `json:"instance_name"`
	KeypairName      *string            `json:"key_pair_name"`
	OSType           *string            `json:"os_type"`
	SecurityGroupID  *string            `json:"security_group_id"`
	VPCID            *string            `json:"vpc_id"`
	CostChargeType   *params.ChargeType `json:"cost_charge_type"`
	HasEIP           *bool              `json:"has_eip"`
	EIPAddress       *string            `json:"eip_address"`
	PrivateIPAddress *string            `json:"private_ip_address"`
	ZoneID           *string            `json:"zone_id"`
	RegionID         *string            `json:"region_id"`
	Status           *string            `json:"status"`
}

// ListInstancesResponse 查询主机列表返回数据
type ListInstancesResponse struct {
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	Total     int        `json:"total"`
	RequestID string     `json:"request_id"`
	Data      []Instance `json:"data"`
}

// ListInstances 查询主机列表
func (s *ECS) ListInstances(p *ListInstancesParams) (r *ListInstancesResponse, err error) {
	err = s.client.Call(http.MethodGet, "/v1/vm/instance", p, nil, &r)
	return
}
