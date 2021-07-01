package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VPCListArgs ...
type VPCListArgs struct {
	Marker *string `json:"marker"`
	Limit  *int    `json:"limit"`
	ZoneID string  `json:"zone_id"`

	VPCIDs []string `json:"vpc_ids"`
}

// VPCListResp ..
type VPCListResp struct {
	common.Response
	Data []*VPC `json:"data"`
}

//VPCList ..
func (d *VPC) VPCList(args *VPCListArgs) (resp *VPCListResp, err error) {
	str := "/api/rio/v1/network/vpc"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str, args.VPCIDs)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
