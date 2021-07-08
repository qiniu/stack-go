package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ImageDeleteArgs 删除镜像参数
type ImageDeleteArgs struct {
	ZoneID                string `json:"header:x-rio-zone-id"`
	ImageID               string `json:"path:id"`
	Force                 *bool  `json:"force,omitempty"`
	DeleteRelatedSnapshot *bool  `json:"delete_related_snapshot,omitempty"`
}

// ImageDeleteResp 删除镜像result
type ImageDeleteResp struct {
	common.Response
}

// ImageDelete 删除镜像
func (d *Image) ImageDelete(args *ImageDeleteArgs) (resp *ImageDeleteResp, err error) {
	url := fmt.Sprintf("%s/image/%s", ComputURLPrefix, args.ImageID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}
