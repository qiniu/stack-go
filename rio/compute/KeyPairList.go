package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairListArgs 查询密钥对列表参数
type KeyPairListArgs struct {
	// 暂不支持分页
	// Marker string `pos:"query:marker"`
	// Limit  int    `pos:"query:limit"`

	ZoneID      string `json:"header:x-rio-zone-id"`
	KeyPairName string `json:"query:keypair_name"`
}

// KeyPairListResp 密钥对列表返回
type KeyPairListResp struct {
	common.Response
	Data []*KeyPairInfo `json:"data"`
}

// KeyPairList 密钥对列表
func (d *KeyPair) KeyPairList(args *KeyPairListArgs) (resp *KeyPairListResp, err error) {
	str := "/api/rio/v1/compute/keypair"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
