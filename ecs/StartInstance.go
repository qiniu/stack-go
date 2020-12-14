package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// StartInstanceParams 主机开机参数
type StartInstanceParams struct {
	RegionID   string `json:"region_id"`
	InstanceID string `json:"instance_id"`

	// 适用于实例规格族d1、i1或者i2等包含本地盘的实例。当d1、i1或者i2的本地盘出现故障时，可通过此参数指定启动实例时，是否将实例恢复到最初的健康状态。取值范围：
	// - true：将实例恢复到最初的健康状态，实例原有本地盘中的数据将会丢失。
	// - false（默认）：不做任何处理，维持现状。
	InitLocalDisk bool `json:"init_local_disk"`
}

// StartInstanceResponse 主机开机返回数据
type StartInstanceResponse struct {
	RequestID string `json:"request_id"`
}

// StartInstance 主机开机
func (s *ECS) StartInstance(p *StartInstanceParams) (resp *StartInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/instance/%s/start", p.InstanceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
