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
func (e *EIP) FindEIPAttachServer(args *FindEIPAttachServerArgs) (resp *FindEIPAttachServerRes, err error) {
	url := fmt.Sprintf("%s/eip/%s/attach_servers", NetworkURLPrefix, args.EIPID)
	req := client.NewRequest(http.MethodGet, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = e.client.Call(req, &resp)
	return
}
