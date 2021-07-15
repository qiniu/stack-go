package compute

import (
	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// NetworkProtocol 网络原型
type NetworkProtocol string

// Direction .
type Direction string

// SecurityGroup 安全组类接口封装
type SecurityGroup struct {
	client             *client.Client
	SecurityGroupID    string               `json:"security_group_id"`
	ZoneID             string               `json:"zone_id"`
	SecurityGroupName  string               `json:"security_group_name"`
	Description        string               `json:"description"`
	SecurityGroupRules []*SecurityGroupRule `json:"security_group_rules"`
}

// SecurityGroupRule 封装
type SecurityGroupRule struct {
	client                *client.Client
	SecurityGroupRuleID   string                       `json:"security_group_rule_id"`
	Direction             common.Direction             `json:"direction"`
	Type                  common.SecurityRuleGrantType `json:"type"`
	RemoteSecurityGroupID *string                      `json:"remote_security_group_id"`
	CIDR                  *string                      `json:"cidr"`
	Protocol              common.NetworkProtocol       `json:"protocol"`
	FromPort              *int                         `json:"from_port"`
	ToPort                *int                         `json:"to_port"`
	Description           string                       `json:"description"`
	CreatedAt             int64                        `json:"created_at"`
}

// NewSecurityGroup 获得实例
func NewSecurityGroup(cli *client.Client) *SecurityGroup {
	return &SecurityGroup{client: cli}
}

// NewSecurityGroupRule 获得实例
func NewSecurityGroupRule(cli *client.Client) *SecurityGroupRule {
	return &SecurityGroupRule{client: cli}
}
