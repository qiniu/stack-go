package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// CreateInstanceParams 主机创建参数
type CreateInstanceParams struct {
	// TODO
}

// CreateInstanceResponse 主机创建返回数据
type CreateInstanceResponse struct {
	// TODO
}

// CreateInstance 主机创建
func (s *ECS) CreateInstance(p *CreateInstanceParams) (resp *CreateInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/instance")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
