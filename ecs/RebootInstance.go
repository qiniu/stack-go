package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// RebootInstanceParams 主机重启参数
type RebootInstanceParams struct {
	RegionID   string `json:"region_id"`
	InstanceID string `json:"instance_id"`

	// 重启ECS实例前，是否强制关机。取值范围：
	// - true：强制关机。相当于典型的断电操作，所有未写入存储设备的缓存数据会丢失。
	// - false（默认）：正常关机。
	ForceStop *bool `json:"force_stop"`
}

// RebootInstanceResponse 主机重启返回数据
type RebootInstanceResponse struct {
	RequestID string `json:"request_id"`
}

// RebootInstance 主机重启
func (s *ECS) RebootInstance(p *RebootInstanceParams) (resp *RebootInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/instance/%s/restart", p.InstanceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
