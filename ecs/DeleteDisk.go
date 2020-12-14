package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteDiskParams 磁盘删除参数
type DeleteDiskParams struct {
	// TODO
}

// DeleteDiskResponse 磁盘删除返回数据
type DeleteDiskResponse struct {
	// TODO
}

// DeleteDisk 磁盘删除
func (s *ECS) DeleteDisk(p *DeleteDiskParams) (resp *DeleteDiskResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/disk/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
