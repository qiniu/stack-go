package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"

	"github.com/qiniu/stack-go/params"
)

// DescribeInstancesParams 查询主机列表参数
type DescribeInstancesParams struct {
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

// DescribeInstancesResponse 查询主机列表返回数据
type DescribeInstancesResponse struct {
	Page      int        `json:"page"`
	Size      int        `json:"size"`
	Total     int        `json:"total"`
	RequestID string     `json:"request_id"`
	Data      []Instance `json:"data"`
}

// DescribeInstances 查询主机列表
func (s *ECS) DescribeInstances(p *DescribeInstancesParams) (resp *DescribeInstancesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/instance").WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
