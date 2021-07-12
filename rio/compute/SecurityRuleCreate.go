package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupRuleCreateArgs 参数
type SecurityGroupRuleCreateArgs struct {
	ZoneID string `json:"zone_id"`

	SecurityGroupID string `json:"security_group_id"`

	Direction             common.Direction             `json:"direction"`
	Type                  common.SecurityRuleGrantType `json:"type"`
	RemoteSecurityGroupID []*string                    `json:"remote_security_group_id"`
	CIDR                  []*string                    `json:"cidr"`
	Protocol              common.NetworkProtocol       `json:"protocol"`
	PortRange             []*string                    `json:"port_range"`
	Description           *string                      `json:"description"`
}

// SecurityGroupRuleCreateResp 返回
type SecurityGroupRuleCreateResp struct {
	common.Response
	Data []*SecurityGroupRule `json:"data"`
}

// SecurityGroupRuleCreate 安全组规则创建
func (d *SecurityGroupRule) SecurityGroupRuleCreate(args *SecurityGroupRuleCreateArgs) (resp *SecurityGroupRuleCreateResp, err error) {
	url := fmt.Sprintf("%s/security_group/rule", ComputURLPrefix)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
