package common

// Response 返回
type Response struct {
	RequestID string  `json:"request_id"`
	Marker    *string `json:"marker,omitempty"`
}

// Direction 方向
type Direction string

// 方向常量
const (
	IngressDirection Direction = "ingress"
	EgressDirection  Direction = "egress"
	AllDirection     Direction = "all"
)
