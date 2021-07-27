package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchCreateArgs 创建交换机参数
type VSwitchCreateArgs struct {
	ZoneID      string  `json:"zone_id"`
	VPCID       string  `json:"vpc_id"`
	VSwitchName *string `json:"vswitch_name"`
	VSwitchCidr string  `json:"vswitch_cidr"`
	Description *string `json:"description"`
	ClientToken *string `json:"client_token"`
}

// VSwitchCreatResp 创建交换机返回
type VSwitchCreatResp struct {
	common.Response
	VSwitchID string `json:"vswitch_id"`
}

// VSwitchCreate 创建交换机
func (v *VSwitch) VSwitchCreate(args *VSwitchCreateArgs) (resp *VSwitchCreatResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/vswitch", NetworkURLPrefix, args.VPCID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
