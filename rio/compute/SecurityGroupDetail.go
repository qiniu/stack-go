package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupDetailArgs 安全组详情参数
type SecurityGroupDetailArgs struct {
	ZoneID          string `json:"zone_id"`
	SecurityGroupID string `json:"security_group_id"`
}

// SecurityGroupDetailResp 安全组详情返回
type SecurityGroupDetailResp struct {
	common.Response
	Data *SecurityGroup `json:"data"`
}

// SecurityGroupDetail 安全组详情
func (s *SecurityGroup) SecurityGroupDetail(args *SecurityGroupDetailArgs) (resp *SecurityGroupDetailResp, err error) {
	url := fmt.Sprintf("%s/security_group/%s/detail", ComputURLPrefix, args.SecurityGroupID)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
