package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// CreateDiskArgs 创建磁盘参数
type CreateDiskArgs struct {
	ZoneID string `json:"zone_id"`

	// 容量大小，以GiB为单位。指定该参数后，其取值必须≥指定快照ID的容量大小。取值范围：
	// - cloud：5~2000
	// - cloud_efficiency：20~32768
	// - cloud_ssd：20~32768
	// - cloud_essd：20~32768
	Size *int64 `json:"size,omitempty"`

	// 云盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 http:// 和 https:// 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	//
	// 默认值：空
	DiskName *string `json:"disk_name,omitempty"`

	// 云盘描述。长度为2~256个英文或中文字符，不能以 http:// 和 https:// 开头。
	//
	// 默认值：空
	Description *string `json:"description,omitempty"`

	// 创建云盘使用的快照。指定该参数后，Size会被忽略，实际创建的云盘大小为指定快照的大小。
	SnapshotID *string `json:"snapshot_id,omitempty"`
}

// CreateDiskResp 创建磁盘返回
type CreateDiskResp struct {
	common.Response
	Data struct {
		DiskID string `json:"disk_id"`
	} `json:"data"`
}

// CreateDisk 创建磁盘
func (d *Disk) CreateDisk(args *CreateDiskArgs) (resp *CreateDiskResp, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/api/proxy/rio/v1/storage/disk")).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
