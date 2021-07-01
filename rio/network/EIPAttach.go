package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//EIPAttachArgs 绑定eip参数
type EIPAttachArgs struct {
	ZoneID           string `json:"zone_id"`
	EIPID            string `json:"eip_id"`
	ServerID         string `json:"server_id"`
	PrivateIPAddress string `json:"private_ip_address"`
}

// EIPAttachResp eip绑定返回
type EIPAttachResp struct {
	*common.Response
}

//EIPAttach 绑定eip
func (d *EIP) EIPAttach(args *EIPAttachArgs) (resp *EIPAttachResp, err error) {
	str := "/api/rio/v1/network/eip"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/attach", args.EIPID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
