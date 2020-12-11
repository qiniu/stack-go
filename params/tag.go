package params

// Tag 标签
type Tag struct {
	ID           string `json:"_id"`
	UID          uint32 `json:"uid"`
	RegionID     string `json:"region_id"`
	ResourceType string `json:"resource_type"`
	ResourceID   string `json:"resource_id"`
	Tag          string `json:"tag"`
}
