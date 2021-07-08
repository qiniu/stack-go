package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnModifyArgs 告警修改参数
type WarnModifyArgs struct {
	ZoneID    string            `pos:"header:x-rio-zone-id"`
	WarnID    *string           `pos:"path:id"`
	WarnName  string            `json:"name"`
	AlertUser string            `json:"alert_user"`
	IsPhone   string            `json:"is_phone"`
	PhoneList string            `json:"phone_list"`
	IsEmail   string            `json:"is_email"`
	EmailList string            `json:"email_list"`
	Enable    string            `json:"enable"`
	Rules     []*ModifyRule     `json:"rules"`
	Instances []*ModifyInstance `json:"instances"`
}

// ModifyRule 规则参数
type ModifyRule struct {
	EventType            common.EventType `json:"event_type"`
	RuleFlag             string           `json:"rule_flag"`
	TimeInterval         int64            `json:"time_interval"`
	Statistice           string           `json:"statistice"`
	ComparisonOperator   string           `json:"comparison_operator"`
	Threshold            int64            `json:"threshold"`
	MediumThreshold      int64            `json:"medium_threshold"`
	SeriousnessThreshold int64            `json:"seriousness_threshold"`
}

// ModifyInstance 实例参数
type ModifyInstance struct {
	InstanceID string `json:"instance_id"`
}

// WarnModifyResp 修改返回
type WarnModifyResp struct {
	common.Response
}

// WarnModify 告警修改
func (d *Warn) WarnModify(args *WarnModifyArgs) (resp *WarnModifyResp, err error) {
	str := "/api/rio/v1/compute/warn"
	req := client.NewRequest(http.MethodPut, fmt.Sprintf(str+"%s", args.WarnID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}