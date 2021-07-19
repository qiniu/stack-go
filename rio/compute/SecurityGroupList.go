package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupListArgs 安全组列表参数
type SecurityGroupListArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupIDs []string `json:"security_group_ids"`
}

// SecurityGroupListResp 安全组列表返回
type SecurityGroupListResp struct {
	common.Response
	Data []*SecurityGroup `json:"data"`
}

// SecurityGroupList 安全组列表
func (s *SecurityGroup) SecurityGroupList(args *SecurityGroupListArgs) (resp *SecurityGroupListResp, err error) {
	url := fmt.Sprintf("%s/security_group", ComputURLPrefix)
	req := client.NewRequest(http.MethodGet, url).WithQueries(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
