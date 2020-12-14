package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteKeypairParams 密钥对删除参数
type DeleteKeypairParams struct {
	// TODO
}

// DeleteKeypairResponse 密钥对删除返回数据
type DeleteKeypairResponse struct {
	// TODO
}

// DeleteKeypair 密钥对删除
func (s *ECS) DeleteKeypair(p *DeleteKeypairParams) (resp *DeleteKeypairResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/keypair")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
