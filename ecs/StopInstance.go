package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// StopInstanceParams 主机关机参数
type StopInstanceParams struct {
	// TODO
}

// StopInstanceResponse 主机关机返回数据
type StopInstanceResponse struct {
	// TODO
}

// StopInstance 主机关机
func (s *ECS) StopInstance(p *StopInstanceParams) (resp *StopInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/instance/:id/stop")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
