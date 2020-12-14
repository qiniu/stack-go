package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// DeleteInstanceParams 主机删除参数
type DeleteInstanceParams struct {
	RegionID   string `json:"region_id"`
	InstanceID string `json:"instance_id"`
}

// DeleteInstanceResponse 主机删除返回数据
type DeleteInstanceResponse struct {
	RequestID string `json:"request_id"`
}

// DeleteInstance 主机删除
func (s *ECS) DeleteInstance(p *DeleteInstanceParams) (resp *DeleteInstanceResponse, err error) {
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf("/v1/vm/instance/%s", p.InstanceID)).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
