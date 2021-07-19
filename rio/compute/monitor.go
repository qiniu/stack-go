package compute

import "github.com/qiniu/stack-go/components/client"

// Monitor 镜像类接口封装
type Monitor struct {
	client *client.Client
}

// NewMonitor 获得镜像实例
func NewMonitor(cli *client.Client) *Monitor {
	return &Monitor{client: cli}
}

// MonitorReturnType 监控信息返回数据类型，默认json
const MonitorReturnType string = "json"

// MonitorDataHistoryBol 监控信息时候查询历史数据，默认1，即默认查询历史数据
const MonitorDataHistoryBol string = "1"

// MonitorMaxCount 监控数据最高返回 200 条，按照供应商返回数据条数，间隔取数
const MonitorMaxCount int = 200

// MonitorDataType 监控数据查询类型
type MonitorDataType string

// 监控数据查询类型枚举
const (
	MonitorDataTypeCPU      MonitorDataType = "cpu"       // CPU
	MonitorDataTypeMem      MonitorDataType = "mem"       // 内存
	MonitorDataTypeDisk     MonitorDataType = "disk"      // 磁盘
	MonitorDataTypeDiskInfo MonitorDataType = "disk_info" // 磁盘（已用空间，使用率等。前提需要安装qga工具）
	MonitorDataTypeNetOut   MonitorDataType = "netout"    // 出网络
	MonitorDataTypeNetIn    MonitorDataType = "netin"     // 入网络
)
