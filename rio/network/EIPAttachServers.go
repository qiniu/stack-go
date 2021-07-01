package network

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
	"github.com/qiniu/stack-go/rio/compute"
)

//FindEIPAttachServerArgs 查询eip可绑定主机列表参数
type FindEIPAttachServerArgs struct {
	ZoneID string `json:"zone_id"`
	EIPID  string `json:"eip_id"`
}

// FindEIPAttachServerRes 查询eip可绑定主机列表返回
type FindEIPAttachServerRes struct {
	common.Response
	Data []*compute.ServerInfo `json:"data"`
}

//FindEIPAttachServer 查询eip可绑定主机列表
func (d *EIP) FindEIPAttachServer(args *FindEIPAttachServerArgs) (resp *FindEIPAttachServerRes, err error) {
	str := "/api/rio/v1/network/eip"
	req := client.NewRequest(http.MethodGet, fmt.Sprintf(str+"/%s/attach_servers", args.EIPID)).WithJSONBody(&args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
