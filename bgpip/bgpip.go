package bgpip

import "github.com/qiniu/stack-go/components/client"

// BGPIP BGP高防
type BGPIP struct {
	client *client.Client
}

// NewBGPIP 获得 BGPIP 实例
func NewBGPIP(cli *client.Client) *BGPIP {
	return &BGPIP{
		client: cli,
	}
}
