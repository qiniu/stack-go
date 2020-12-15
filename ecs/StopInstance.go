package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// StopInstanceParams 主机关机参数
type StopInstanceParams struct {
	RegionID   string `json:"region_id"`
	InstanceID string `json:"instance_id"`

	// 停止按量付费ECS实例后，是否继续计费。取值：
	// - StopCharging：停止计费。有关StopCharging生效的条件，请参见按量付费实例停机不收费的停机不收费开启条件章节。
	// - KeepCharging：继续计费。
	// 默认值：如果您在ECS控制台上开启了VPC类型实例默认停机不收费（详情请参见开启按量付费实例停机不收费），并符合开启条件，则默认为StopCharging。否则，默认为KeepCharging。
	StoppedMode InstanceStoppedMode `json:"stopped_mode"`

	// 停止实例时的是否强制关机策略。取值范围：
	// - true：强制关机。
	// - false（默认）：正常关机流程
	ForceStop *bool `json:"force_stop"`

	// 是否确认关机。仅对i1型实例规格族生效，且为i1型的实例规格族的必选参数。默认值：false。
	ConfirmStop *bool `json:"confirm_stop"`
}

// StopInstanceResponse 主机关机返回数据
type StopInstanceResponse struct {
	RequestID string `json:"request_id"`
}

// StopInstance 主机关机
func (s *ECS) StopInstance(p *StopInstanceParams) (resp *StopInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/instance/%s/stop", p.InstanceID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
