package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairDetachArgs 密钥对解绑参数
type KeyPairDetachArgs struct {
	ZoneID      string `json:"zone_id"`
	KeyPairName string `json:"keypair_name"`
	ServerID    string `json:"server_id"`
}

// KeyPairDetachResp 返回
type KeyPairDetachResp struct {
	common.Response
}

// KeyPairDetach 密钥对解绑
func (k *KeyPair) KeyPairDetach(args *KeyPairDetachArgs) (resp *KeyPairDetachResp, err error) {
	url := fmt.Sprintf("%s/keypair/%s/detach", ComputURLPrefix, args.KeyPairName)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = k.client.Call(req, &resp)
	return
}
