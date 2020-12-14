package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DescribeDiskParams 磁盘详情参数
type DescribeDiskParams struct {
	// TODO
}

// DescribeDiskResponse 磁盘详情返回数据
type DescribeDiskResponse struct {
	// TODO
}

// DescribeDisk 磁盘详情
func (s *ECS) DescribeDisk(p *DescribeDiskParams) (resp *DescribeDiskResponse, err error) {
	req := client.NewRequest(http.MethodGet, "/v1/vm/disk/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
