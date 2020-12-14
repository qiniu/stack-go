package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteInstanceParams 主机删除参数
type DeleteInstanceParams struct {
	// TODO
}

// DeleteInstanceResponse 主机删除返回数据
type DeleteInstanceResponse struct {
	// TODO
}

// DeleteInstance 主机删除
func (s *ECS) DeleteInstance(p *DeleteInstanceParams) (resp *DeleteInstanceResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/instance/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
