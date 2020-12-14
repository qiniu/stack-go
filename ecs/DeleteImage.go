package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteImageParams 镜像删除参数
type DeleteImageParams struct {
	// TODO
}

// DeleteImageResponse 镜像删除返回数据
type DeleteImageResponse struct {
	// TODO
}

// DeleteImage 镜像删除
func (s *ECS) DeleteImage(p *DeleteImageParams) (resp *DeleteImageResponse, err error) {
	req := client.NewRequest(http.MethodDelete, "/v1/vm/image/:id")
	// .WithQueries(p).WithRegionID(p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
