package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupListArgs 安全组列表参数
type SecurityGroupListArgs struct {
	// Marker string `pos:"query:marker"`
	// Limit  int    `pos:"query:limit"` 暂不支持分页
	ZoneID string `json:"zone_id"`

	SecurityGroupIDs []string `json:"security_group_ids"`
}

// SecurityGroupListResp 安全组列表返回
type SecurityGroupListResp struct {
	common.Response
	Data []*SecurityGroup `json:"data"`
}

// SecurityGroupList 安全组列表
func (d *SecurityGroup) SecurityGroupList(args *SecurityGroupListArgs) (resp *SecurityGroupListResp, err error) {
	str := "/api/rio/v1/compute/security_group"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str)).WithQueries(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
