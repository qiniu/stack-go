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
func (e *EIP) EIPModify(args *EIPModifyArgs) (resp *EIPModifyResp, err error) {
	url := fmt.Sprintf("%s/eip/%s", NetworkURLPrefix, args.EIPID)
	req := client.NewRequest(http.MethodPut, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = e.client.Call(req, &resp)
	return
}
