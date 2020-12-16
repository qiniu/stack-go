package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"

	"github.com/qiniu/stack-go/params"
)

// DescribeAvailableResourceParams 查询可用资源
type DescribeAvailableResourceParams struct {

	// 目标地域ID。您可以调用DescribeRegions查看最新的七牛云地域列表。
	RegionID string `json:"region_id"`

	// 要查询的资源类型。取值范围：
	// - Zone：可用区
	// - IoOptimized：I/O优化
	// - InstanceType：实例规格
	// - SystemDisk：系统盘
	// - DataDisk：数据盘
	// - Network：网络类型
	// - ddh：专有宿主机
	// 参数DestinationResource的取值方式请参见上文接口说明。
	DestinationResource params.ResourceType `json:"destination_resource"`

	// 可用区ID。
	//
	// 默认值：无，表示随机分配当前地域下的可用区，返回该地域（RegionId）下所有可用区的符合查询条件的资源。
	ZoneID *string `json:"zone_id,omitempty"`

	// 资源的计费方式。更多详情，请参见计费概述。取值范围：
	// - PrePaid：包年包月
	// - PostPaid：按量付费
	// 默认值：PostPaid
	InstanceChargeType *params.ChargeType `json:"instance_charge_type,omitempty"`

	// 实例规格。更多详情，请参见实例规格族，也可以调用DescribeInstanceTypes接口获得最新的规格表。
	//
	// 当参数DestinationResource取值为SystemDisk或者DataDisk时，InstanceType为必需参数。
	InstanceType *string `json:"instance_type,omitempty"`

	// 系统盘类型。取值范围：
	// 若参数DestinationResource取值为SystemDisk、InstanceType或者DataDisk时，参数SystemDiskCategory不是必需参数。
	//
	// 默认值：cloud_efficiency
	SystemDiskCategory *DiskCategory `json:"system_disk_category,omitempty"`

	// 数据盘类型。取值范围：
	DataDiskCategory *DiskCategory `json:"data_disk_category,omitempty"`

	// 网络类型。取值范围：
	// - vpc：专有网络
	// - classic：经典网络
	NetworkCategory *NetworkType `json:"network_category,omitempty"`
}

// DescribeAvailableResourceResponse 查询可用资源列表返回数据
type DescribeAvailableResourceResponse struct {
	RequestID string `json:"request_id"`
	Data      struct {
		AvailableZone []*AvailableZoneType `json:"available_zone,omitempty"`
	} `json:"data"`
}

// DescribeAvailableResource 查询可用资源列表
func (s *ECS) DescribeAvailableResource(p *DescribeAvailableResourceParams) (resp *DescribeAvailableResourceResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/resource_available").WithJSONBody(p).WithRegionID(&p.RegionID)
	err = s.client.Call(req, &resp)
	return
}
