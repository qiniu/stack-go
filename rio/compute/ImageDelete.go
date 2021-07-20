package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ImageDeleteArgs 删除镜像参数
type ImageDeleteArgs struct {
	ZoneID                string `json:"zone_id"`
	ImageID               string `json:"image_id"`
	Force                 *bool  `json:"force,omitempty"`
	DeleteRelatedSnapshot *bool  `json:"delete_related_snapshot,omitempty"`
}

// ImageDeleteResp 删除镜像返回
type ImageDeleteResp struct {
	common.Response
}

// ImageDelete 删除镜像
func (i *Image) ImageDelete(args *ImageDeleteArgs) (resp *ImageDeleteResp, err error) {
	url := fmt.Sprintf("%s/image/%s", ComputURLPrefix, args.ImageID)
	req := client.NewRequest(http.MethodDelete, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = i.client.Call(req, &resp)
	return
}
