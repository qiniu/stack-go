package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VPCModifyArgs vpc修改参数
type VPCModifyArgs struct {
	ZoneID      string  `json:"zone_id"`
	VPCName     *string `json:"vpc_name"`
	Description *string `json:"description"`
	VPCID       string  `json:"vpc_id"`
}

// VPCModifyResp vpc修改返回
type VPCModifyResp struct {
	common.Response
}

// VPCModify vpc修改
func (v *VPC) VPCModify(args *VPCModifyArgs) (resp *VPCModifyResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/modify", NetworkURLPrefix, args.VPCID)
	req := client.NewRequest(http.MethodPut, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
