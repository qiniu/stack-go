package common

// Response 返回
type Response struct {
	RequestID string  `json:"request_id"`
	Marker    *string `json:"marker,omitempty"`
}
