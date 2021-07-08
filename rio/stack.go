package rio

import (
	"github.com/qiniu/stack-go/components/auth"
	"github.com/qiniu/stack-go/components/log"
	"github.com/qiniu/stack-go/rio/compute"
	"github.com/qiniu/stack-go/rio/storage"

	"github.com/qiniu/stack-go/components/client"
)

// SDK 版本
const version = "0.0.2"

// Stack 七牛云主机服务 SDK
type Stack struct {
	log    log.Logger
	client *client.Client
}

// New 构造 SDK 实例
func New(log log.Logger, endpoint string, credential *auth.Credential) *Stack {
	return &Stack{
		log:    log,
		client: client.New(log, endpoint, credential),
	}
}

// Version 获取当前 SDK 版本
func (*Stack) Version() string {
	return version
}

// Disk 获取 Disk 管理对象
func (s *Stack) Disk() *storage.Disk {
	return storage.NewDisk(s.client)
}

// Image 获取 Image管理对象
func (s *Stack) Image() *compute.Image {
	return compute.NewImage(s.client)
}

// KeyPair 密钥对
func (s *Stack) KeyPair() *compute.KeyPair {
	return compute.NewKeyPair(s.client)
}

// Monitor ..
func (s *Stack) Monitor() *compute.Monitor {
	return compute.NewMonitor(s.client)
}

// SecurityGroup ..
func (s *Stack) SecurityGroup() *compute.SecurityGroup {
	return compute.NewSecurityGroup(s.client)
}

// SecurityGroupRule ..
func (s *Stack) SecurityGroupRule() *compute.SecurityGroupRule {
	return compute.NewSecurityGroupRule(s.client)
}

// Server ..
func (s *Stack) Server() *compute.Server {
	return compute.NewServer(s.client)
}

// Warn ..
func (s *Stack) Warn() *compute.Warn {
	return compute.NewWarn(s.client)
}
