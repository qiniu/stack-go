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

// EIPCreateResp eip创建返回
type EIPCreateResp struct {
	common.Response
}

// EIPCreate 创建eip
func (e *EIP) EIPCreate(args *EIPCreateArgs) (resp *EIPCreateResp, err error) {
	url := fmt.Sprintf("%s/eip", NetworkURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = e.client.Call(req, &resp)
	return
}
