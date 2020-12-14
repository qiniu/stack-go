package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// RenewInstanceParams 主机续费参数
type RenewInstanceParams struct {
	// TODO
}

// RenewInstanceResponse 主机续费返回数据
type RenewInstanceResponse struct {
	// TODO
}

// RenewInstance 主机续费
func (s *ECS) RenewInstance(p *RenewInstanceParams) (resp *RenewInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/instance/:id/renew")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
