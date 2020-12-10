package params

import "time"

// AutoRenew 自动续费
type AutoRenew struct {
	ID           string       `json:"_id"`
	UID          uint32       `json:"uid"`
	RegionID     string       `json:"region_id"`
	ResourceType ResourceType `json:"resource_type"`
	ResourceID   string       `json:"resource_id"`
	CostParams   *CostParams  `json:"cost_params"`
	UpdatedAt    time.Time    `json:"updated_at"`
	CreatedAt    time.Time    `json:"created_at"`
}
