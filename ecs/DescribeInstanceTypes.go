package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeInstanceTypesParams 主机规格列表参数
type DescribeInstanceTypesParams struct {
	Family *string `json:"family"`
}

// DescribeInstanceTypesResponse 主机规格列表返回数据
type DescribeInstanceTypesResponse struct {
	RequestID string         `json:"request_id"`
	Data      []InstanceType `json:"data"`
}

// DescribeInstanceTypes 主机规格列表
func (s *ECS) DescribeInstanceTypes(p *DescribeInstanceTypesParams) (resp *DescribeInstanceTypesResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/instance_types").WithQueries(p)
	err = s.client.Call(req, &resp)
	return
}
