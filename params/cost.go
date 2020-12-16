package params

// CostParams 计费参数
type CostParams struct {
	CostChargeType ChargeType `json:"cost_charge_type"` // 付费方式
	CostChargeMode ChargeMode `json:"cost_charge_mode"` // 计费方式
	CostPeriodUnit PeriodUnit `json:"cost_period_unit"` // 计费单位
	CostPeriod     int        `json:"cost_period"`      // 计费周期
}

// PeriodUnit 资源计费单位
type PeriodUnit string

// 小时、日、周、月、年
const (
	PeriodUnitOnHour    PeriodUnit = "Hour"
	PeriodUnitOnDay     PeriodUnit = "Day"
	PeriodUnitOnWeekly  PeriodUnit = "Week"
	PeriodUnitOnMonthly PeriodUnit = "Month"
	PeriodUnitOnYearly  PeriodUnit = "Year"
)

// ChargeMode 资源计费方式: 按...计费
type ChargeMode string

// 计费方式
const (
	ChargeModeBandwidth        ChargeMode = "PayByBandwidth"        // 按照带宽计费
	ChargeModeTraffic          ChargeMode = "PayByTraffic"          // 按照流量计费
	ChargeModeDisk             ChargeMode = "PayByDisk"             // 按磁盘使用计费
	ChargeModeInstance         ChargeMode = "PayByInstance"         // 按主机规格使用计费
	ChargeModeSlbSpec          ChargeMode = "PayBySlbSpec"          // 按SLB实例规格计费
	ChargeModeRedis            ChargeMode = "PayByRedis"            // 按 Redis 实例规格计费
	ChargeModeMongo            ChargeMode = "PayByMongo"            // 按 Mongo 实例规格计费
	ChargeModeVRouterInterface ChargeMode = "PayByVRouterInterface" // 按 VRouterInterface 实例规格计费
	ChargeModePayBy95          ChargeMode = "PayBy95"               // 按增强型95计费
	ChargeModeRds              ChargeMode = "PayByRds"              // 按 RDS 实例规格计费
	ChargeModeVpn              ChargeMode = "PayByVpn"              // 按 vpn 规格计费
	ChargeModeNATGatewaySpec   ChargeMode = "PayByNATGatewaySpec"   // 按 高速通道 规格计费
	ChargeModeNewBgpIP         ChargeMode = "PayByNewBgpIp"         // 安全产品新BGP
	ChargeModeBgpIP            ChargeMode = "PayByBgpIp"            // 安全产品高防
	ChargeModeWaf              ChargeMode = "PayByWaf"              // 安全产品WAF
	ChargeModePolarDB          ChargeMode = "PayByPolarDB"          // 按 PolarDB 节点规格计费
)

// ChargeType 支付类型
type ChargeType string

// 按需计费（按量使用，月度账单）、预付费（包年包月，订单支付后使用）
const (
	PostPaid ChargeType = "PostPaid" // 按需计费
	PrePaid  ChargeType = "PrePaid"  // 预付费
)

// CostInfo 付费信息
type CostInfo struct {
	CostChargeType ChargeType `json:"cost_charge_type"` // 付费方式
	CostChargeMode ChargeMode `json:"cost_charge_mode"` // 按什么计费
}
