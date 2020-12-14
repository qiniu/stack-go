package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// AttachKeypairParams 密钥对绑定参数
type AttachKeypairParams struct {
	// TODO
}

// AttachKeypairResponse 密钥对绑定返回数据
type AttachKeypairResponse struct {
	// TODO
}

// AttachKeypair 密钥对绑定
func (s *ECS) AttachKeypair(p *AttachKeypairParams) (resp *AttachKeypairResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/keypair/:id/attach")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
