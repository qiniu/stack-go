package rio

import (
	"github.com/qiniu/stack-go/components/cookieauth"
	"github.com/qiniu/stack-go/components/log"
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
func New(log log.Logger, cfg *cookieauth.Config) *Stack {
	return &Stack{
		log:    log,
		client: client.NewWithTransport(log, cfg.Endpoint, cookieauth.NewTransport(cfg)),
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
