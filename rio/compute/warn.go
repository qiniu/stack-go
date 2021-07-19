package compute

import (
	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// Warn 告警
type Warn struct {
	client *client.Client
}

// WarnStatus 规则组状态
type WarnStatus string

// 规则组状态 Enum
const (
	WarnStatusOK               WarnStatus = "OK"                // 正常
	WarnStatusInsufficientData WarnStatus = "INSUFFICIENT_DATA" // 数据不足
	WarnStatusALARM            WarnStatus = "ALARM"             // 报警中
)

// NewWarn 获得实例
func NewWarn(cli *client.Client) *Warn {
	return &Warn{client: cli}
}

// InstanceType 告警资源类型
type InstanceType string

// 告警资源类型 Enum
const (
	WarnInstanceCloudHost InstanceType = "cloud" // 云主机
)

// WarnAlertLevel 告警管理告警级别枚举
type WarnAlertLevel string

// 告警管理告警级别枚举
// LOST|CRITICAL|WARNING|OK_RECOVERY|OK
const (
	WarnAlertLevelLOST       WarnAlertLevel = "LOST"        // 紧急告警
	WarnAlertLevelCRITICAL   WarnAlertLevel = "CRITICAL"    // 严重告警
	WarnAlertLevelWARNING    WarnAlertLevel = "WARNING"     // 中度告警
	WarnAlertLevelOKRECOVERY WarnAlertLevel = "OK_RECOVERY" // 告警恢复
	WarnAlertLevelOK         WarnAlertLevel = "OK"          // 正常
)

// AlertMessageDetail 告警信息详情
type AlertMessageDetail struct {
	Confirm         string                 `json:"confirm"`             // 确认状态，0：未确认，1：已确认
	AlertID         string                 `json:"alert_id"`            // 告警信息ID
	AllInfo         string                 `json:"all_info_cn"`         // 提示信息
	AlertContent    string                 `json:"alert_content_cn"`    // 告警详情
	AlertLevel      common.WarnAlertLevel  `json:"alert_status"`        // 告警级别
	ResourceID      string                 `json:"resource_id"`         // 实例 ID
	ResourceScName  string                 `json:"resource_sc_name_cn"` // 告警实例类型
	ResourceName    string                 `json:"resource_name_cn"`    // 实例名
	ResourceAddress string                 `json:"resource_address"`    // 实例所处 IP
	InsertTime      int64                  `json:"insert_time"`         // 告警信息产生时间
	UpdateTime      int64                  `json:"update_time"`         // 更新时间，可根据创建时间计算持续时间
	HistoryAlerts   []*HistoryAlertMessage `json:"history_alerts"`      // 历史告警信息
}

// HistoryAlertMessage 历史告警数据
type HistoryAlertMessage struct {
	Confirm         string                `json:"confirm"`             // 确认状态，0：未确认，1：已确认
	AlertID         string                `json:"alert_id"`            // 告警信息ID
	AllInfo         string                `json:"all_info_cn"`         // 提示信息
	AlertContent    string                `json:"alert_content_cn"`    // 告警详情
	AlertLevel      common.WarnAlertLevel `json:"level"`               // 告警级别
	ResourceID      string                `json:"resource_id"`         // 实例 ID
	ResourceScName  string                `json:"resource_sc_name_cn"` // 告警实例类型
	ResourceName    string                `json:"resource_name_cn"`    // 实例名
	ResourceAddress string                `json:"resource_address"`    // 实例所处 IP
	InsertTime      int64                 `json:"insert_time"`         // 告警信息产生时间
	UpdateTime      int64                 `json:"update_time"`         // 更新时间，可根据创建时间计算持续时间
}
