package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairCreateArgs 密钥对创建参数
type KeyPairCreateArgs struct {
	ZoneID        string `json:"zone_id"`
	KeyPairName   string `json:"key_pair_name"`
	PublicKeyBody string `json:"public_key_body"`
}

// KeyPairCreateResp 密钥对创建返回
type KeyPairCreateResp struct {
	common.Response
	Data struct {
		*KeyPairInfo
		PrivateKeyBody *string `json:"private_key_body"`
	} `json:"data"`
}

// KeyPairCreate 密钥对创建
func (d *KeyPair) KeyPairCreate(args *KeyPairCreateArgs) (resp *KeyPairCreateResp, err error) {
	str := "/api/rio/v1"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/compute/keypair")).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
