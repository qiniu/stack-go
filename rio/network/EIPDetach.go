package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// EIPDetachArgs 解绑eip参数
type EIPDetachArgs struct {
	ZoneID   string `json:"zone_id"`
	EIPID    string `json:"eip_id"`
	ServerID string `json:"server_id"`
}

// EIPDetachResp 解绑eip返回
type EIPDetachResp struct {
	common.Response
}

// EIPDetach eip解绑返回
func (e *EIP) EIPDetach(args *EIPDetachArgs) (resp *EIPDetachResp, err error) {
	url := fmt.Sprintf("%s/eip/%s/detach", NetworkURLPrefix, args.EIPID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = e.client.Call(req, &resp)
	return
}
