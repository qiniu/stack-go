package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DetachKeypairParams 密钥对解绑参数
type DetachKeypairParams struct {
	// TODO
}

// DetachKeypairResponse 密钥对解绑返回数据
type DetachKeypairResponse struct {
	// TODO
}

// DetachKeypair 密钥对解绑
func (s *ECS) DetachKeypair(p *DetachKeypairParams) (resp *DetachKeypairResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/keypair/:id/attach")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
