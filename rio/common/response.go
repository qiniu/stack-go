package common

type Response struct {
	RequestID string  `json:"request_id"`
	Marker    *string `json:"marker,omitempty"`
}
