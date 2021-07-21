package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ServerReplaceSystemDiskArgs 主机重置系统盘参数
type ServerReplaceSystemDiskArgs struct {
	ZoneID   string `json:"zone_id"`
	ServerID string `json:"server_id"`
	Size     int64  `json:"size"`
	ImageID  string `json:"image_id"`
}

// ServerReplaceSystemDiskResp 主机重置系统盘返回
type ServerReplaceSystemDiskResp struct {
	common.Response
}

// ServerReplaceSystemDisk 主机重置系统盘
func (s *Server) ServerReplaceSystemDisk(args *ServerReplaceSystemDiskArgs) (resp *ServerReplaceSystemDiskResp, err error) {
	url := fmt.Sprintf("%s/server/%s/replace_system_disk", ComputURLPrefix, args.ServerID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = s.client.Call(req, &resp)
	return
}
