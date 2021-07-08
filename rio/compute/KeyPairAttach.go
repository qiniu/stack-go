package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairAttachArgs 密钥对绑定参数
type KeyPairAttachArgs struct {
	ZoneID      string `json:"zone_id"`
	KeyPairName string `json:"keypair_name"`
	ServerID    string `json:"server_id"`
}

// KeyPairAttachResp 密钥对返回
type KeyPairAttachResp struct {
	common.Response
}

// KeyPairAttach 密钥对绑定
func (d *KeyPair) KeyPairAttach(args *KeyPairAttachArgs) (resp *KeyPairAttachResp, err error) {
	str := "/api/rio/v1/compute/keypair"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/attach", args.KeyPairName)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
