package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//EIPDetachArgs 解绑eip参数
type EIPDetachArgs struct {
	ZoneID   string `pos:"header:x-rio-zone-id"`
	EIPID    string `pos:"path:id"`
	ServerID string `json:"server_id"`
}

// EIPDetachResp 解绑IP返回
type EIPDetachResp struct {
	common.Response
}

//EIPDetach eip解绑返回
func (d *EIP) EIPDetach(args *EIPDetachArgs) (resp *EIPDetachResp, err error) {
	str := "/api/rio/v1/network/eip"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/detach", args.EIPID)).WithJSONBody(&args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
