package common

// SecurityRuleGrantType 安全组授权类型
type SecurityRuleGrantType int

const (
	SecurityGroupType SecurityRuleGrantType = 1
	CIDRType          SecurityRuleGrantType = 2
)
