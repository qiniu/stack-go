package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VPCListArgs 列表参数
type VPCListArgs struct {
	Marker *string `json:"marker"`
	Limit  *int    `json:"limit"`
	ZoneID string  `json:"zone_id"`

	VPCIDs []string `json:"vpc_ids"`
}

// VPCListResp 列表返回
type VPCListResp struct {
	common.Response
	Data []*VPC `json:"data"`
}

//VPCList vpc列表
func (v *VPC) VPCList(args *VPCListArgs) (resp *VPCListResp, err error) {
	url := fmt.Sprintf("%s/vpc", NetworkURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
