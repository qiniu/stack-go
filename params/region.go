package params

// Region 地域
type Region struct {
	RegionID  string  `json:"region_id"`
	LocalName string  `json:"local_name"`
	Zones     []*Zone `json:"zones"`
}

// Zone 可用区
type Zone struct {
	RegionID  string `json:"region_id"`
	ZoneID    string `json:"zone_id"`
	LocalName string `json:"local_name"`
}
