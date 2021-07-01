package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//EIPModifyArgs EIP修改参数
type EIPModifyArgs struct {
	ZoneID    string `json:"zone_id"`
	EIPID     string `json:"eip_id"`
	EIPName   string `json:"eip_name,omitempty"`
	Bandwidth uint   `json:"bandwidth,omitempty"`
}

// EIPModifyResp eip修改返回
type EIPModifyResp struct {
	common.Response
}

//EIPModify EIP修改
func (d *EIP) EIPModify(args *EIPModifyArgs) (resp *EIPModifyResp, err error) {
	str := "/api/rio/v1/network/eip"
	req := client.NewRequest(http.MethodPut, fmt.Sprintf(str+"/%s", args.EIPID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
