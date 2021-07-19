package compute

import (
	"fmt"
	"net/http"

	"github.com/qiniu/stack-go/components/client"
	"github.com/qiniu/stack-go/rio/common"
)

// ImageImportArgs 镜像导入参数
type ImageImportArgs struct {
	ZoneID                 string  `json:"zone-id"`
	ImgaeID                string  `json:"image_id"`
	SourceURL              string  `json:"source_url"`
	DestinationZoneID      string  `json:"destination_zone_id"`
	DestinationImageName   *string `json:"destination_image_name,omitempty"`
	DestinationDescription *string `json:"destination_description,omitempty"`
}

// ImageImportResp 自定义导入镜像返回
type ImageImportResp struct {
	common.Response
}

// ImageImport 镜像导入
func (i *Image) ImageImport(args *ImageImportArgs) (resp *ImageImportResp, err error) {
	url := fmt.Sprintf("%s/image/%s", ComputURLPrefix, args.ImgaeID)
	req := client.NewRequest(http.MethodPost, url).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = i.client.Call(req, &resp)
	return
}
