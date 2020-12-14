package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// ImportKeypairParams 密钥对导入参数
type ImportKeypairParams struct {
	// TODO
}

// ImportKeypairResponse 密钥对导入返回数据
type ImportKeypairResponse struct {
	// TODO
}

// ImportKeypair 密钥对导入
func (s *ECS) ImportKeypair(p *ImportKeypairParams) (resp *ImportKeypairResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/keypair")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
