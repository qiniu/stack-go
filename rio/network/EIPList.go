package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// EIPListArgs eip列表参数
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

// EIPListResp eip列表返回
type EIPListResp struct {
	common.Response
	Data []*EIPInfo `json:"data"`
}

// EIPInfo eip信息
type EIPInfo struct {
	ZoneID     string `json:"zone_id"`
	ServerID   string `json:"server_id"`
	ServerName string `json:"server_name"`
	EIPID      string `json:"eip_id"`
	EIPName    string `json:"eip_name"`
	EIPAddress string `json:"eip_address"`
	Bandwidth  uint   `json:"bandwidth"`
	CreatedAt  int64  `json:"created_at"`
}

// EIPList eip列表
func (e *EIP) EIPList(args *EIPListArgs) (resp *EIPListResp, err error) {
	url := fmt.Sprintf("%s/eip", NetworkURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = e.client.Call(req, &resp)
	return
}
