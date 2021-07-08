package compute

import "github.com/qiniu/stack-go/components/client"

// Image 镜像类接口封装
type Image struct {
	client *client.Client
}

// NewImage 获得镜像实例
func NewImage(cli *client.Client) *Image {
	return &Image{client: cli}
}
