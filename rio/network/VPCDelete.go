package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VPCDeleteArgs 删除参数
type VPCDeleteArgs struct {
	ZoneID string `json:"zone_id"`
	VPCID  string `json:"vpc_id"`
}

// VPCDeleteResp 删除返回
type VPCDeleteResp struct {
	common.Response
}

// VPCDelete vpc删除
func (v *VPC) VPCDelete(args *VPCDeleteArgs) (resp *VPCCreateResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/delete", NetworkURLPrefix, args.VPCID)
	req := client.NewRequest(http.MethodDelete, url).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
