package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchModifyArgs 交换机修改参数
type VSwitchModifyArgs struct {
	ZoneID      string  `json:"header:x-rio-zone-id"`
	VPCID       string  `json:"path:id"`
	VSwitchID   string  `json:"path:vswitch_id"`
	VSwitchName string  `json:"vswitch_name"`
	Description *string `json:"description"`
}

// VSwitchModifyResp 交换机修改返回
type VSwitchModifyResp struct {
	common.Response
}

// VSwitchModify 交换机修改
func (v *VSwitch) VSwitchModify(args *VSwitchModifyArgs) (resp *VPCModifyResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/vswitch/%s/modify", NetworkURLPrefix, args.VPCID, args.VSwitchID)
	req := client.NewRequest(http.MethodPut, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
