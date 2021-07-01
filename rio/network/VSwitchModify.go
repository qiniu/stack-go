package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchModifyArgs ..
type VSwitchModifyArgs struct {
	ZoneID      string  `json:"header:x-rio-zone-id"`
	VPCID       string  `json:"path:id"`
	VSwitchID   string  `json:"path:vswitch_id"`
	VSwitchName string  `json:"vswitch_name"`
	Description *string `json:"description"`
}

// VSwitchModifyResp ..
type VSwitchModifyResp struct {
	common.Response
}

//VSwitchModify ..
func (d *VSwitch) VSwitchModify(args *VSwitchModifyArgs) (resp *VPCModifyResp, err error) {
	str := "/api/rio/v1/network/vpc"
	req := client.NewRequest(http.MethodPut, fmt.Sprintf(str+"/%s/vswitch/%s/modify", args.VPCID, args.VSwitchID)).WithJSONBody(&args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
