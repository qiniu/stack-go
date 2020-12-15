package ecs

import (
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/params"
)

// CreateDiskParams 磁盘创建参数
type CreateDiskParams struct {
	// 所属的地域ID。您可以调用DescribeRegions查看最新的七牛云地域列表。
	RegionID string `json:"region_id"`

	// 在指定可用区内创建一块按量付费云盘。
	// - 如果您不设置InstanceId，则ZoneId为必填参数。
	// - 您不能同时指定ZoneId和InstanceId。
	ZoneID string `json:"zone_id"`

	// 容量大小，以GiB为单位。指定该参数后，其取值必须≥指定快照ID的容量大小。取值范围：
	// - cloud：5~2000
	// - cloud_efficiency：20~32768
	// - cloud_ssd：20~32768
	// - cloud_essd：20~32768
	Size int64 `json:"size"` // GB

	// 云盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 默认值：空
	DiskName *string `json:"disk_name"`

	// 云盘描述。长度为2~256个英文或中文字符，不能以 http:// 和 https:// 开头。
	//
	// 默认值：空
	Description *string `json:"description"`

	// 是否加密云盘。默认值：false
	Encrypted *bool `json:"encrypted"`

	// 数据盘的云盘种类。取值范围：
	// - cloud：普通云盘
	// - cloud_efficiency：高效云盘
	// - cloud_ssd：SSD云盘
	// - cloud_essd：ESSD云盘
	// 默认值：cloud
	DiskCategory *DiskCategory `json:"disk_category"`

	// 创建一块ESSD云盘时，设置云盘的性能等级。取值范围：
	// - PL1（默认）：单盘最高随机读写IOPS 5万
	// - PL2：单盘最高随机读写IOPS 10万
	// - PL3：单盘最高随机读写IOPS 100万
	// 有关如何选择ESSD性能等级，请参见ESSD云盘。
	PerformanceLevel *ESSDPerformanceLevel `json:"performance_level"`

	// 创建云盘使用的快照。指定该参数后，Size会被忽略，实际创建的云盘大小为指定快照的大小。2013年7月15日及以前的快照不能用来创建云盘。
	SnapshotID *string `json:"snapshot_id"`

	// 创建一块包年包月云盘，并自动挂载到指定的包年包月实例（InstanceId）上。
	// - 设置实例ID后，会忽略您设置的ResourceGroupId、Tag.N.Key、Tag.N.Value、ClientToken和KMSKeyId参数。
	// - 您不能同时指定ZoneId和InstanceId。
	// 默认值：空，代表创建的是按量付费云盘，云盘所属地由RegionId和ZoneId确定。
	InstanceID *string `json:"instance_id"`

	// 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
	ClientToken *string `json:"client_token"`

	*params.CostParams
}

// CreateDiskResponse 磁盘创建返回数据
type CreateDiskResponse struct {
	RequestID string `json:"request_id"`
	Data      *struct {
		DiskID    *string `json:"disk_id"`    // 按量创建返回 ID
		OrderHash *string `json:"order_hash"` // 预付费创建返回订单号
	} `json:"data"`
}

// CreateDisk 磁盘创建
func (s *ECS) CreateDisk(p *CreateDiskParams) (resp *CreateDiskResponse, err error) {
	req := client.NewRequest(http.MethodPost, "/v1/vm/disk").WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
