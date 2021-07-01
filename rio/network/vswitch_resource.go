package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchResourceArgs ..
type VSwitchResourceArgs struct {
	ZoneID    string  `pos:"header:x-rio-zone-id"`
	VPCID     *string `pos:"path:id,required"`
	VSwitchID *string `pos:"path:vswitch_id,required"`
}

// VSwitchResourceResp ..
type VSwitchResourceResp struct {
	common.Response
	Data []VSwitchResourceInfo `json:"data"`
}

//VSwitchResourceInfo  ..
type VSwitchResourceInfo struct {
	ResourceID   string `json:"resource_id"`
	ResourceName string `json:"resource_name"`
	ResourceType string `json:"resource_type"`
}

//VSwitchResource ..
func (d *VSwitch) VSwitchResource(args *VSwitchResourceArgs) (resp *VSwitchResourceResp, err error) {
	str := "/api/rio/v1/network/vpc"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/%s/vswitch/%s/resource", args.VPCID, args.VSwitchID)).WithJSONBody(&args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
