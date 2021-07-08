package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// SecurityGroupRuleCreateArgs ..
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

// SecurityGroupRuleCreateResp ..
type SecurityGroupRuleCreateResp struct {
	common.Response
	Data []*SecurityGroupRule `json:"data"`
}

// SecurityGroupRuleCreate ..
func (d *SecurityGroupRule) SecurityGroupRuleCreate(args *SecurityGroupRuleCreateArgs) (resp *SecurityGroupRuleCreateResp, err error) {
	str := "/api/rio/v1/compute/security_group"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s/rule", args.SecurityGroupID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
