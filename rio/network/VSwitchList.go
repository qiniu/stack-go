package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchListArgs 交换机列表参数
type VSwitchListArgs struct {
	Marker *string `json:"marker"`
	Limit  *int    `json:"limit"`
	ZoneID string  `json:"zone_id"`

	VPCID      string   `json:"id"`
	VSwitchIDs []string `json:"vswitch_ids"`
}

// VSwitchListResp 返回
type VSwitchListResp struct {
	common.Response
	Data []*VSwitch `json:"data"`
}

// VSwitchList 交换机列表
func (v *VSwitch) VSwitchList(args *VSwitchListArgs) (resp *VSwitchListResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/vswitch", NetworkURLPrefix, args.VPCID)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = v.client.Call(req, &resp)
	return
}