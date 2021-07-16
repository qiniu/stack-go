package storage

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ModifyDiskArgs 修改磁盘参数
type ModifyDiskArgs struct {
	ZoneID string `json:"zone_id"`

	DiskID string `json:"disk_id"`

	// 磁盘名称。长度为2~128个英文或中文字符。必须以大小字母或中文开头，不能以 `http://` 和 `https://` 开头。可以包含数字、半角冒号（:）、下划线（_）或者连字符（-）。
	DiskName *string `json:"disk_name,omitempty"`

	// 磁盘描述。 长度为2~256个英文或中文字符，不能以 `http://` 和 `https://` 开头。
	Description *string `json:"description,omitempty"`

	// 磁盘是否随实例释放。默认值：无，无表示不改变当前的值。
	DeleteWithServer *bool `json:"delete_with_server,omitempty"`
}

// ModifyDiskResp 修改磁盘返回
type ModifyDiskResp struct {
	common.Response
}

//ModifyDisk 修改磁盘
func (d *Disk) ModifyDisk(args *ModifyDiskArgs) (resp *ModifyDiskResp, err error) {
	url := fmt.Sprintf("%s/disk/%s", StorageURLPrefix, args.DiskID)
	req := client.NewRequest(http.MethodPatch, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
