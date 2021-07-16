package common

type Response struct {
	RequestID string  `json:"request_id"`
	Marker    *string `json:"marker,omitempty"`
}
type Direction string

const (
	IngressDirection Direction = "ingress"
	EgressDirection  Direction = "egress"
	AllDirection     Direction = "all"
)
