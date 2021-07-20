package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// EIPDeleteArgs 删除eip参数
type EIPDeleteArgs struct {
	ZoneID string `json:"zone_id"`
	EIPID  string `json:"eip_id"`
}

// EIPDeleteResp eip删除返回
type EIPDeleteResp struct {
	common.Response
}

// EIPDelete eip删除
func (e *EIP) EIPDelete(args *EIPDeleteArgs) (resp *EIPDeleteResp, err error) {
	url := fmt.Sprintf("%s/eip/%s", NetworkURLPrefix, args.EIPID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = e.client.Call(req, &resp)
	return
}
