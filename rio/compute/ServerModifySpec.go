package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerModifySpecArgs 参数
type ServerModifySpecArgs struct {
	ZoneID      string `json:"zone_id"`
	ServerID    string `json:"server_id"`              // 实例ID。
	FlavorID    string `json:"flavor_id,omitempty"`    // 实例的目标规格。更多详情，请参见实例规格族，也可以调用DescribeServerTypes接口获得最新的规格表。
	ClientToken string `json:"client_token,omitempty"` // 保证请求幂等性。从您的客户端生成一个参数值，确保不同请求间该参数值唯一。
}

// ServerModifySpecResp 返回
type ServerModifySpecResp struct {
	common.Response
}

// ServerModifySpec 主机修改
func (d *Server) ServerModifySpec(args *ServerModifySpecArgs) (resp *ServerModifySpecResp, err error) {
	url := fmt.Sprintf("%s/server/%s/spec", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPatch, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
