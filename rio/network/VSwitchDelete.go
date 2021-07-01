package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//VSwitchDeleteArgs 删除参数
type VSwitchDeleteArgs struct {
	ZoneID    string `json:"zone_id"`
	VPCID     string `json:"vpc_id"`
	VSwitchID string `json:"vswitch_id"`
}

// VSwitchDeleteResp ...
type VSwitchDeleteResp struct {
	common.Response
}

//VSwitchDelete 删除
func (d *VSwitch) VSwitchDelete(args *VSwitchDeleteArgs) (resp *VSwitchDeleteResp, err error) {
	str := "/api/rio/v1/network/vpc"
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf(str+"/%s/vswitch/%s/delete", args.VPCID, args.VSwitchID)).WithJSONBody(&args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
