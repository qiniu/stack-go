package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// RebootInstanceParams 主机重启参数
type RebootInstanceParams struct {
	// TODO
}

// RebootInstanceResponse 主机重启返回数据
type RebootInstanceResponse struct {
	// TODO
}

// RebootInstance 主机重启
func (s *ECS) RebootInstance(p *RebootInstanceParams) (resp *RebootInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/instance/:id/restart")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
