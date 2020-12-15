package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// DescribeRegionsResponse 查询地域列表返回数据
type DescribeRegionsResponse struct {
	Data []params.Region `json:"data"`
}

// DescribeRegions 查询地域列表
func (s *ECS) DescribeRegions() (r *DescribeRegionsResponse, err error) {
	err = s.client.Call(client.NewRequest(http.MethodGet, "/v1/cm/region"), &r)
	return
}
