package stack

import (
	"github.com/qiniu/stack-go/bgpip"
	"github.com/qiniu/stack-go/components/auth"
	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/components/log"

	"github.com/qiniu/stack-go/bss"
	"github.com/qiniu/stack-go/ecs"
	"github.com/qiniu/stack-go/vpc"
)

// SDK 版本
const version = "0.0.1"

// Stack 七牛云主机服务 SDK
type Stack struct {
	log        log.Logger
	credential *auth.Credential
	client     *client.Client
}

// New 构造 SDK 实例
func New(log log.Logger, endpoint string, credential *auth.Credential) *Stack {
	return &Stack{
		log:        log,
		credential: credential,
		client:     client.New(log, endpoint, credential),
	}
}

// Version 获取当前 SDK 版本
func (*Stack) Version() string {
	return version
}

// ECS 获取 ECS 管理对象
func (s *Stack) ECS() *ecs.ECS {
	return ecs.NewECS(s.client)
}

// VPC 获取 VPC 管理对象
func (s *Stack) VPC() *vpc.VPC {
	return vpc.NewVPC(s.client)
}

// BSS 获取 BSS 管理对象
func (s *Stack) BSS() *bss.BSS {
	return bss.NewBSS(s.client)
}

// BGPIP 获取 BGPIP 管理对象
func (s *Stack) BGPIP() *bgpip.BGPIP {
	return bgpip.NewBGPIP(s.client)
}
