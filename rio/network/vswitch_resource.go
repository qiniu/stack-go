package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchResourceArgs 交换机下资源列表参数
type VSwitchResourceArgs struct {
	ZoneID    string `json:"zone_id"`
	VPCID     string `json:"id"`
	VSwitchID string `json:"vswitch_id"`
}

// VSwitchResourceResp 交换机下资源列表返回
type VSwitchResourceResp struct {
	common.Response
	Data []VSwitchResourceInfo `json:"data"`
}

// VSwitchResourceInfo 资源信息
type VSwitchResourceInfo struct {
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

// VSwitchResource 交换机下资源列表
func (v *VSwitch) VSwitchResource(args *VSwitchResourceArgs) (resp *VSwitchResourceResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/vswitch/%s/resource", NetworkURLPrefix, args.VPCID, args.VSwitchID)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
