package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//EIPDeleteArgs 删除eip参数
type EIPDeleteArgs struct {
	ZoneID string `json:"zone_id"`
	EIPID  string `json:"eip_id"`
}

// EIPDeleteResp eip删除返回
type EIPDeleteResp struct {
	common.Response
}

//EIPDelete eip删除
func (d *EIP) EIPDelete(args *EIPDeleteArgs) (resp *EIPDeleteResp, err error) {
	str := "/api/rio/v1/network/eip"

	req := client.NewRequest(http.MethodDelete, fmt.Sprintf(str+"/%s", args.EIPID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
