package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchResourceArgs 参数
type VSwitchResourceArgs struct {
	ZoneID    string `json:"zone_id"`
	VPCID     string `json:"path:id,required"`
	VSwitchID string `json:"path:vswitch_id,required"`
}

// VSwitchResourceResp 返回
type VSwitchResourceResp struct {
	common.Response
	Data []VSwitchResourceInfo `json:"data"`
}

//VSwitchResourceInfo  信息
type VSwitchResourceInfo struct {
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

//VSwitchResource 交换机
func (d *VSwitch) VSwitchResource(args *VSwitchResourceArgs) (resp *VSwitchResourceResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/vswitch/%s/resource", NetworkURLPrefix, args.VPCID, args.VSwitchID)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
