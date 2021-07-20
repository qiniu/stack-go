package compute

import "github.com/qiniu/stack-go/components/client"

// MonitorDataCommonInfo 监控数据节点时间
type MonitorDataCommonInfo struct {
	TimeStamp int64 `json:"time_stamp"`
}

// Monitor 镜像类接口封装
type Monitor struct {
	client *client.Client
}

// NewMonitor 获得镜像实例
func NewMonitor(cli *client.Client) *Monitor {
	return &Monitor{client: cli}
}

// ServerMonitorDataType 云主机监控详细数据
type ServerMonitorDataType struct {
	CPU           *float64 `json:"cpu"`             // 实例vCPU的使用比例，单位：百分比（%）
	NetInBytes    *float64 `json:"net_in_bytes"`    // 接收宽带，单位：Byte/s
	NetInPackets  *float64 `json:"net_in_packets"`  // 实例在TimeStamp时刻接收数据包数，单位：个
	NetInErrs     *float64 `json:"net_in_errs"`     // 设备驱动器检测到的接收错误包数量，单位：个
	NetOutBytes   *float64 `json:"net_out_bytes"`   // 发送宽带，单位：Byte/s
	NetOutPackets *float64 `json:"net_out_packets"` // 实例在TimeStamp时刻发送数据包数，单位：个
	NetOutErrs    *float64 `json:"net_out_errs"`    // 设备驱动器检测到的发送错误包数量，单位：个
	IOPSRead      *float64 `json:"io_ps_read"`      // 实例系统盘I/O读操作，单位：次/s
	IOPSWrite     *float64 `json:"io_ps_write"`     // 实例系统盘I/O写操作，单位：次/s
	BPSRead       *float64 `json:"bps_read"`        // 实例系统盘读带宽，单位：Byte/s
	BPSWrite      *float64 `json:"bps_write"`       // 实例系统盘写带宽，单位：Byte/s

	// ===== 暂无返回
	IntranetRX        *float64 `json:"intranet_rx"`        // 实例在TimeStamp时刻接收的内网数据流量，单位：kbits
	IntranetTX        *float64 `json:"intranet_tx"`        // 实例在TimeStamp时刻发送的内网数据流量，单位：kbits
	IntranetBandwidth float64  `json:"intranet_bandwidth"` // 实例的内网带宽，单位时间内的网络流量，单位：kbits/s
	InternetRX        float64  `json:"internet_rx"`        // 实例在TimeStamp时刻接收的公网数据流量，单位：kbits
	InternetTX        float64  `json:"internet_tx"`        // 实例在TimeStamp时刻发送的公网数据流量，单位：kbits
	InternetBandwidth float64  `json:"internet_bandwidth"` // 实例的公网带宽，单位时间内的网络流量，单位：kbits/s
	// =====
}

// ServerMonitorDataInfo 云主机监控数据
type ServerMonitorDataInfo struct {
	MonitorDataCommonInfo
	ServerMonitorDataType
}

// MonitorReturnType 返回类型
const MonitorReturnType string = "json"

// MonitorDataHistoryBol 数据历史
const MonitorDataHistoryBol string = "1"

// MonitorMaxCount 最大值
const MonitorMaxCount int = 200

// MonitorDataType 数据类型
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

// DiskMonitorDataInfo 磁盘监控数据
type DiskMonitorDataInfo struct {
	MonitorDataCommonInfo
	DiskMonitorDataType
}

// DiskMonitorDataType 磁盘监控详细数据
type DiskMonitorDataType struct {
	IOPSRead  *float64 `json:"io_ps_read"`  // 系统盘I/O读操作，单位：次/s。
	IOPSWrite *float64 `json:"io_ps_write"` // 系统盘I/O写操作，单位：次/s。
	IOPSTotal *float64 `json:"io_ps_total"` // 系统盘I/O读写总操作，单位：次/s。
	BPSRead   *float64 `json:"bps_read"`    // 系统盘读带宽，单位：Byte/s。
	BPSWrite  *float64 `json:"bps_write"`   // 系统盘写带宽，单位：Byte/s。
	BPSTotal  *float64 `json:"bps_total"`   // 系统盘读写总带宽，单位：Byte/s。
}

// EipMonitorDataInfo 网卡监控数据
type EipMonitorDataInfo struct {
	MonitorDataCommonInfo
	EipMonitorDataType
}

// EipMonitorDataType 网卡监控详细数据
type EipMonitorDataType struct {
	EipRX        float64 `json:"eip_rx"`        // 流入的带宽。单位为 Byte/s
	EipTX        float64 `json:"eip_tx"`        // 流出的带宽。单位为 Byte/s
	EipFlow      float64 `json:"eip_flow"`      // 流入和流出的带宽总和 Byte/s
	EipBandwidth float64 `json:"eip_bandwidth"` // 带宽值，该值等于EipFlow/60，单位为 Byte/s
	EipPackets   float64 `json:"eip_packets"`   // 包数量
	EipErrs      float64 `json:"eip_errs"`      // 错误包数量
	EipDrop      float64 `json:"eip_drop"`      // 丢失包数量
}
