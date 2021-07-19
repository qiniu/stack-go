package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

//EIPListArgs EIP列表参数
type EIPListArgs struct {
	Marker string `pos:"query:marker"`
	Limit  int    `pos:"query:limit"`

	ZoneID     string            `json:"zone_id"`
	Status     *common.EIPStatus `json:"status"`
	ServerID   *string           `json:"server_id"`
	EIPName    *string           `json:"eip_name"`
	EIPID      *string           `json:"eip_id"`
	EIPAddress *string           `json:"eip_address"`
}

// EIPListResp EIP列表返回
type EIPListResp struct {
	common.Response
	Data []*EIPInfo `json:"data"`
}

//EIPList EIP列表
func (e *EIP) EIPList(args *EIPListArgs) (resp *EIPListResp, err error) {
	url := fmt.Sprintf("%s/eip", NetworkURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = e.client.Call(req, &resp)
	return
}