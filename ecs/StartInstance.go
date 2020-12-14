package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// StartInstanceParams 主机开机参数
type StartInstanceParams struct {
	// TODO
}

// StartInstanceResponse 主机开机返回数据
type StartInstanceResponse struct {
	// TODO
}

// StartInstance 主机开机
func (s *ECS) StartInstance(p *StartInstanceParams) (resp *StartInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/instance/:id/start")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
