package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// VSwitchListArgs 交换机列表参数
type VSwitchListArgs struct {
	Marker *string `json:"query:marker"`
	Limit  *int    `json:"query:limit"`
	ZoneID string  `json:"header:x-rio-zone-id"`

	VPCID      string   `json:"path:id"`
	VSwitchIDs []string `json:"query:vswitch_ids"`
}

// VSwitchListResp 返回
type VSwitchListResp struct {
	common.Response
	Data []*VSwitch `json:"data"`
}

//VSwitchList 交换机列表
func (d *VSwitch) VSwitchList(args *VSwitchListArgs) (resp *VSwitchListResp, err error) {
	url := fmt.Sprintf("%s/vpc/%s/vswitch", NetworkURLPrefix, args.VPCID)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
