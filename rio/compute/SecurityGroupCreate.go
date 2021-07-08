package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupCreateArgs 创建安全组参数
type SecurityGroupCreateArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupName *string `json:"security_group_name"`
	Description       *string `json:"description"`
	ClientToken       *string `json:"client_token"`
}

// SecurityGroupCreateResp 创建安全组返回
type SecurityGroupCreateResp struct {
	common.Response
	Data *SecurityGroup `json:"data"`
}

// SecurityGroupCreate 创建安全组
func (d *SecurityGroup) SecurityGroupCreate(args *SecurityGroupCreateArgs) (resp *SecurityGroupCreateResp, err error) {
	url := fmt.Sprintf("%s/security_group", ComputURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
