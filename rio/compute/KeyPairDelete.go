package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairDeleteArgs  删除密钥对参数
type KeyPairDeleteArgs struct {
	ZoneID      string `json:"zone_id"`
	KeyPairName string `json:"keypair_name"`
}

// KeyPairDeleteResp 删除密钥对返回
type KeyPairDeleteResp struct {
	common.Response
}

//KeyPairDelete 删除密钥对
func (d *KeyPair) KeyPairDelete(args *KeyPairDeleteArgs) (resp *KeyPairDeleteResp, err error) {
	str := "/api/rio/v1/compute/keypair"
	req := client.NewRequest(http.MethodDelete, fmt.Sprintf(str+"/%s", args.KeyPairName)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
