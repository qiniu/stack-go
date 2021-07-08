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

// ImageImportResp 自定义导入镜像result
type ImageImportResp struct {
	common.Response
}

// ImageImport 镜像导入
func (d *Image) ImageImport(args *ImageImportArgs) (resp *ImageImportResp, err error) {
	str := "/api/rio/v1/compute/image"
	req := client.NewRequest(http.MethodPost, fmt.Sprintf(str+"/%s", args.ImgaeID)).WithJSONBody(args).WithZoneID(&args.ZoneID)
	err = d.client.Call(req, &resp)
	return
}