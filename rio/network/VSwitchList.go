package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchListArgs ..
type VSwitchListArgs struct {
	Marker *string `json:"query:marker"`
	Limit  *int    `json:"query:limit"`
	ZoneID string  `json:"header:x-rio-zone-id"`

	VPCID      *string  `json:"path:id"`
	VSwitchIDs []string `json:"query:vswitch_ids"`
}

// VSwitchListResp ..
type VSwitchListResp struct {
	common.Response
	Data []*VSwitch `json:"data"`
}

//VSwitchList ..
func (d *VSwitch) VSwitchList(args *VSwitchListArgs) (resp *VSwitchListResp, err error) {
	str := "/api/rio/v1/network/vpc"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/%s/vswitch", args.VPCID)).WithJSONBody(&args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
