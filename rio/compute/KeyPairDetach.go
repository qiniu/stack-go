package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// KeyPairDetachArgs ..
type KeyPairDetachArgs struct {
	ZoneID      string `json:"zone_id"`
	KeyPairName string `json:"keypair_name"`
	ServerID    string `json:"server_id"`
}

// KeyPairDetachResp ..
type KeyPairDetachResp struct {
	common.Response
}

// KeyPairDetach ..
func (d *KeyPair) KeyPairDetach(args *KeyPairDetachArgs) (resp *KeyPairDetachResp, err error) {
	str := "/api/rio/v1/compute/keypair"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/detach", args.KeyPairName)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
