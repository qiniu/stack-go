package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VPCCreateArgs 创建参数
type VPCCreateArgs struct {
	ZoneID      string  `json:"zone_id"`
	CIDR        *string `json:"cidr"`
	VPCName     *string `json:"vpc_name"`
	Description *string `json:"description"`
	ClientToken *string `json:"client_token"`
}

// VPCCreateResp 返回
type VPCCreateResp struct {
	common.Response
	VPCID string `json:"vpc_id"`
}

// CreateVPCWithVSwitchArgs 绑定交换机参数
type CreateVPCWithVSwitchArgs struct {
	ZoneID  string             `json:"zone_id"`
	VPC     *VPCCreateArgs     `json:"vpc"`
	VSwitch *VSwitchCreateArgs `json:"vswitch"`
}

// VPCCreate VPC创建
func (v *VPC) VPCCreate(args *CreateVPCWithVSwitchArgs) (resp *VPCCreateResp, err error) {
	url := fmt.Sprintf("%s/vpc", NetworkURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
