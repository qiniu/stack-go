package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeDisksParams 磁盘列表参数
type DescribeDisksParams struct {
	// TODO
}

// DescribeDisksResponse 磁盘列表返回数据
type DescribeDisksResponse struct {
	// TODO
}

// DescribeDisks 磁盘列表
func (s *ECS) DescribeDisks(p *DescribeDisksParams) (resp *DescribeDisksResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/disk")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
