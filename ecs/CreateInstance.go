package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// CreateInstanceParams 主机创建参数
type CreateInstanceParams struct {
	*params.CostParams

	RegionID        string  `json:"region_id,omitempty"`         // 实例所属的地域 ID。您可以调用 DescribeRegions 查看最新的七牛云地域列表。
	ImageID         string  `json:"image_id,omitempty"`          // 镜像 ID，启动实例时选择的镜像资源。您可以通过 DescribeImages 查询您可以使用的镜像资源。
	InstanceType    string  `json:"instance_type,omitempty"`     // 实例的资源规格。
	SecurityGroupID string  `json:"security_group_id,omitempty"` // 指定新创建实例所属于的安全组 ID。同一个安全组内的实例之间可以互相访问，一个安全组能容纳的实例数量视安全组类型而定，具体请参见使用限制的安全组章节。
	ZoneID          *string `json:"zone_id,omitempty"`           // 实例所属的可用区ID，您可以调用DescribeZones获取可用区列表。默认值：系统随机选择。

	// 系统盘的云盘种类。取值范围：
	//
	// - cloud_efficiency：高效云盘
	// - cloud_ssd：SSD云盘
	// - cloud：普通云盘
	// - ephemeral_ssd：本地SSD盘
	//
	// 已停售的实例规格且非 I/O 优化实例默认值为 cloud，否则默认值为 cloud_efficiency。
	SystemDiskCategory *DiskCategory `json:"system_disk_category,omitempty"`

	// 系统盘大小，单位为GiB。取值范围：20~500
	// 该参数的取值必须大于或者等于max{20, ImageSize}。
	// 默认值：max{40, ImageSize}
	SystemDiskSize *int64 `json:"system_disk_size,omitempty"`

	// 系统盘名称。长度为 2~128 个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	SystemDiskName *string `json:"system_disk_name,omitempty"`

	// 系统盘描述。长度为 2~256 个英文或中文字符，不能以 http:// 和 https:// 开头。
	SystemDiskDescription *string `json:"system_disk_description,omitempty"`

	// 创建 ESSD 云盘作为系统盘使用时，设置云盘的性能等级。
	//
	// 取值范围：
	// - PL1（默认）：单盘最高随机读写 IOPS 5万
	// - PL2：单盘最高随机读写 IOPS 10万
	// - PL3：单盘最高随机读写 IOPS 100万。
	//
	// 有关如何选择 ESSD 性能等级，请参见 ESSD 云盘。
	SystemDiskPerformanceLevel *ESSDPerformanceLevel `json:"system_disk_performance_level,omitempty"`

	DataDisk []*DiskInfo `json:"data_disk,omitempty"`

	// 虚拟交换机ID。如果您创建的是VPC类型ECS实例，必须指定虚拟交换机ID，且安全组和虚拟交换机在同一个专有网络VPC中。
	VSwitchID *string `json:"v_switch_id,omitempty"`

	// 实例的名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。
	// 可以包含数字、半角冒号（:）、下划线（_）、英文句号（.）或者连字符（-）。
	// 如果没有指定该参数，默认值为实例的InstanceId。
	InstanceName *string `json:"instance_name,omitempty"`

	// 云服务器的主机名。
	// - 英文句号（.）和短横线（-）不能作为首尾字符，更不能连续使用。
	// - Windows实例：字符长度为2~15，不支持英文句号（.），不能全是数字。允许大小写英文字母、数字和短横线（-）。
	// - 其他类型实例（Linux等）：字符长度为2~64，支持多个英文句号（.），英文句号之间为一段，每段允许大小写英文字母、数字和短横线（-）。
	HostName *string `json:"host_name,omitempty"`

	// 实例的描述。长度为 2~256 个英文或中文字符，不能以 http:// 和 https:// 开头。
	Description *string `json:"description,omitempty"`

	// 实例的密码。长度为8至30个字符，必须同时包含大小写英文字母、数字和特殊符号中的三类字符。特殊符号可以是：
	//
	// ```
	// ()`~!@#$%^&*-_+=|{}[]:;'<>,.?/
	// ```
	// 其中，Windows实例不能以斜线号（/）为密码首字符。
	//
	// @GSE:NOTE: 如果传入Password参数，建议您使用HTTPS协议发送请求，避免密码泄露。
	Password *string `json:"password,omitempty"`

	// 是否使用镜像预设的密码。使用该参数时，Password参数必须为空，同时您需要确保使用的镜像已经设置了密码。
	PasswordInherit *bool `json:"password_inherit,omitempty"`

	// 密钥对名称。
	// - Windows实例，忽略该参数。默认为空。即使填写了该参数，仍旧只执行Password的内容。
	// - Linux实例的密码登录方式会被初始化成禁止。为提高实例安全性，强烈建议您使用密钥对的连接方式。
	KeyPairName *string `json:"key_pair_name,omitempty"`

	// 是否开启安全加固。取值范围：
	// - Active：启用安全加固，只对系统镜像生效。
	// - Deactive：不启用安全加固，对所有镜像类型生效。
	SecurityEnhancementStrategy *SecurityEnhancementStrategy `json:"security_enhancement_strategy,omitempty"`

	VpcID            *string `json:"vpc_id,omitempty"`       // VPC ID
	ClientToken      *string `json:"client_token,omitempty"` // 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	EnableAssignIPv6 bool    `json:"enable_assign_ipv6"`     // 是否免费分配ipv6地址,默认不分配，需要先检查实例是否支持ipv6

	EipInfo               *CreateInstanceEip `json:"eip_info,omitempty,"`               // 弹性公网 IP 信息
	EnableAutoRenew       *bool              `json:"enable_auto_renew,omitempty"`       // 是否打开自动续费
	CustomizedReleaseTime *string            `json:"customized_release_time,omitempty"` // 后付费定时删除的时间
	AutoReleaseTime       *string            `json:"auto_release_time"`                 // 自动释放时间

}

// CreateInstanceResponse 主机创建返回数据
type CreateInstanceResponse struct {
	RequestID string `json:"request_id"`
	Data      *struct {
		InstanceID *string `json:"instance_id"` // 按量创建返回 ID
		OrderHash  *string `json:"order_hash"`  // 预付费创建返回订单号
	} `json:"data"`
}

// CreateInstance 主机创建
func (s *ECS) CreateInstance(p *CreateInstanceParams) (resp *CreateInstanceResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/instance").WithJSONBody(p).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
