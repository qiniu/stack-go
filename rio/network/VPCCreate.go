package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//VPCCreateArgs 创建参数
type VPCCreateArgs struct {
	ZoneID      string  `json:"zone_id"`
	CIDR        *string `json:"cidr"`
	VPCName     *string `json:"vpc_name"`
	Description *string `json:"description"`
	ClientToken *string `json:"client_token"`
}

// VPCCreateResp ..
type VPCCreateResp struct {
	common.Response
	VPCID string `json:"vpc_id"`
}

// CreateVPCWithVSwitchArgs ..
type CreateVPCWithVSwitchArgs struct {
	ZoneID  string             `json:"zone_id"`
	VPC     *VPCCreateArgs     `json:"vpc"`
	VSwitch *VSwitchCreateArgs `json:"vswitch"`
}

//VPCCreate VPC创建
func (d *VPC) VPCCreate(args *CreateVPCWithVSwitchArgs) (resp *VPCCreateResp, err error) {
	str := "/api/rio/v1/network/vpc"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
