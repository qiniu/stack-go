package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairListArgs 查询密钥对列表参数
type KeyPairListArgs struct {
	ZoneID      string `json:"zone_id"`
	KeyPairName string `json:"keypair_name"`
}

// KeyPairListResp 密钥对列表返回
type KeyPairListResp struct {
	common.Response
	Data []*KeyPairInfo `json:"data"`
}

// KeyPairList 密钥对列表
func (k *KeyPair) KeyPairList(args *KeyPairListArgs) (resp *KeyPairListResp, err error) {
	url := fmt.Sprintf("%s/keypair", ComputURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = k.client.Call(req, &resp)
	return
}
