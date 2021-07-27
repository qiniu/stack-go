package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VPCListArgs vpc列表参数
type VPCListArgs struct {
	Marker *string `json:"marker"`
	Limit  *int    `json:"limit"`
	ZoneID string  `json:"zone_id"`

	VPCIDs []string `json:"vpc_ids"`
}

// VPCListResp vpc列表返回
type VPCListResp struct {
	common.Response
	Data []*VPCInfo `json:"data"`
}

// VPCInfo VPC信息
type VPCInfo struct {
	VPCID       string           `json:"vpc_id"`
	ZoneID      string           `json:"zone_id"`
	Status      common.VPCStatus `json:"status"`
	VPCName     string           `json:"vpc_name"`
	VSwitchIDs  []string         `json:"vswitch_ids"`
	CIDR        string           `json:"cidr"`
	Description string           `json:"description"`
	CreatedAt   int64            `json:"created_at"`
}

// VPCList vpc列表
func (v *VPC) VPCList(args *VPCListArgs) (resp *VPCListResp, err error) {
	url := fmt.Sprintf("%s/vpc", NetworkURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}
