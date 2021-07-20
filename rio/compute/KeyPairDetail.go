package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairDetailArgs 查看密钥对详情参数
type KeyPairDetailArgs struct {
	ZoneID      string `json:"zone_id"`
	KeyPairName string `json:"keypair_name"`
}

// KeyPairDetailResp  密钥对详情返回
type KeyPairDetailResp struct {
	common.Response
	Data *KeyPairInfo `json:"data"`
}

// KeyPairDetail 密钥对详情
func (k *KeyPair) KeyPairDetail(args *KeyPairDetailArgs) (resp *KeyPairDetailResp, err error) {
	url := fmt.Sprintf("%s/keypair/%s", ComputURLPrefix, args.KeyPairName)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = k.client.Call(req, &resp)
	return
}