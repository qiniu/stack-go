package compute

import (
	"github.com/qiniu/stack-go/components/client"
)

// Server 主机
type Server struct {
	client *client.Client
}

// NewServer 获得主机实例
func NewServer(cli *client.Client) *Server {
	return &Server{client: cli}
}
