package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchDeleteArgs 删除交换机参数
type VSwitchDeleteArgs struct {
	ZoneID    string `json:"zone_id"`
	VPCID     string `json:"vpc_id"`
	VSwitchID string `json:"vswitch_id"`
}

// VSwitchDeleteResp 删除交换机返回
type VSwitchDeleteResp struct {
	common.Response
}

// VSwitchDelete 删除交换机
func (v *VSwitch) VSwitchDelete(args *VSwitchDeleteArgs) (resp *VSwitchDeleteResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/vswitch/%s/delete", NetworkURLPrefix, args.VPCID, args.VSwitchID)
	req := client.NewRequest(http.MethodDelete, url).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
