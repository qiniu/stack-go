package params

import "time"

// AutoRenew 自动续费
type AutoRenew struct {
	RegionID     string       `json:"region_id"`     // EIP所在的地域
	ResourceType ResourceType `json:"resource_type"` // 资源类别
	ResourceID   string       `json:"resource_id"`   // 资源 ID
	CostParams   *CostParams  `json:"cost_params"`   // 计费参数
	UpdatedAt    time.Time    `json:"updated_at"`    // 创建时间
	CreatedAt    time.Time    `json:"created_at"`    // 更新时间
}
