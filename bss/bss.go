package bss

import "github.com/qiniu/stack-go/components/client"

// BSS 商务类接口封装
type BSS struct {
	client *client.Client
}

// NewBSS 获得 BSS 实例
func NewBSS(cli *client.Client) *BSS {
	return &BSS{
		client: cli,
	}
}
