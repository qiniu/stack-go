package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VPCDeleteArgs ..
type VPCDeleteArgs struct {
	ZoneID string `json:"zone_id"`
	VPCID  string `json:"vpc_id"`
}

// VPCDeleteResp ..
type VPCDeleteResp struct {
	common.Response
}

//VPCDelete vpc删除
func (d *VPC) VPCDelete(args *VPCDeleteArgs) (resp *VPCCreateResp, err error) {
	str := "/api/rio/v1/network/vpc"
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf(str+"/%s/delete", args.VPCID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
