package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// EIPCreateArgs 创建eip参数
type EIPCreateArgs struct {
	ZoneID      string  `json:"zone_id"`
	EIPName     *string `json:"eip_name,omitempty"`
	Bandwidth   uint    `json:"bandwidth"`
	ClientToken *string `json:"client_token"`
}

// EIPCreateResp EIP创建返回
type EIPCreateResp struct {
	common.Response
}

// EIPCreate 创建eip
func (d *EIP) EIPCreate(args *EIPCreateArgs) (resp *EIPCreateResp, err error) {
	str := "/api/rio/v1/network/eip"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
