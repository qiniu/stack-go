package ecs

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
)

// AttachDiskParams 磁盘挂载参数
type AttachDiskParams struct {
	RegionID string `json:"region_id"`
	DiskID   string `json:"disk_id"`

	// 目标ECS实例的ID。
	InstanceID string `json:"instance_id"`

	// 释放实例时，该云盘是否随实例一起释放。
	// - true：释放。
	// - false：不释放。云盘会转换成按量付费数据盘而被保留下来。
	// 默认值：false
	//
	// @GSD:NOTE 说明 将 DeleteWithInstance 置为false后，一旦ECS实例被安全控制，即OperationLocks中标记了"LockReason" : "security"，释放ECS实例时会忽略云盘的该属性，被同时释放。
	DeleteWithInstance *bool `json:"delete_with_instance"`

	// 是否作为系统盘挂载。 默认值：false
	//
	// @GSD:NOTE 说明 设置为 `Bootable=true` 时，目标ECS实例必须处于无系统盘状态。
	Bootable *bool `json:"bootable,omitempty"`
}

// AttachDiskResponse 磁盘挂载返回数据
type AttachDiskResponse struct {
	RequestID string `json:"request_id"`
}

// AttachDisk 磁盘挂载
func (s *ECS) AttachDisk(p *AttachDiskParams) (resp *AttachDiskResponse, err error) {
	req := client.NewRequest(http.MethodPost, fmt.Sprintf("/v1/vm/disk/%s/attach", p.DiskID)).WithRegionID(&p.RegionID).WithJSONBody(p)
	err = s.client.Call(req, &resp)
	return
}
