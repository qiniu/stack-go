package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// WarnCreateArgs 告警创建参数
type WarnCreateArgs struct {
	ZoneID     string `json:"zone_id"`
	WarnDetail struct {
		Name      string            `json:"name"`          // 名称
		AlertUser string            `json:"alert_user"`    // 告警联系人
		IsPhone   string            `json:"is_phone"`      // 是否启用短信通知，1：启用，0：禁用
		PhoneList string            `json:"phone_list"`    // 电话号，使用英文","隔开
		IsEmail   string            `json:"is_Email"`      // 是否启用邮件通知，1：启用，0：禁用
		EmailList string            `json:"email_list"`    // 邮箱，使用英文","隔开
		Rules     []*CreateRule     `json:"rules"`         // 告警规则
		Instances []*CreateInstance `json:"instance_list"` // 告警实例
	} `json:"warn_detail"`
}

// CreateInstance 告警实例创建参数
type CreateInstance struct {
	InstanceID string `json:"instance_id"`
}

// CreateRule 告警规则创建参数
type CreateRule struct {
	EventType            common.EventType `json:"rule_type"`
	MediumThreshold      int64            `json:"medium_threshold"`
	SeriousnessThreshold int64            `json:"seriousness_threshold"`
	TimeInterval         int64            `json:"time_interval"`
	Statistics           string           `json:"statistics"`
	ComparisonOperator   string           `json:"comparison_operator"`
	Threshold            int64            `json:"threshold"`
}

// WarnCreateResp 创建告警返回
type WarnCreateResp struct {
	common.Response
}

// WarnCreate 告警创建
func (d *Warn) WarnCreate(args *WarnCreateArgs) (resp *WarnCreateResp, err error) {
	url := fmt.Sprintf("%s/warn", ComputURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
