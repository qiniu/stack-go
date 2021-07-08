package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ImageCreateArgs 创建镜像参数
type ImageCreateArgs struct {
	ZoneID       string  `json:"zone_id"`
	SnapshotID   string  `json:"snapshot_id"`
	ServerID     string  `json:"server_id"`
	ImageName    *string `json:"image_name"`
	ImageVersion *string `json:"image_version"`
	Description  *string `json:"description"`
	ClientToken  *string `json:"client_token"`
	DiskFormat   string  `json:"disk_format"` // 镜像格式
}

// ImageCreateResp 创建镜像返回
type ImageCreateResp struct {
	common.Response
}

// ImageCreate 创建快照
func (d *Image) ImageCreate(args *ImageCreateArgs) (resp *ImageCreateResp, err error) {
	str := "/api/rio/v1"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/compute/image")).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
