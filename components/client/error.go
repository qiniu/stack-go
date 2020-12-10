package client

import "fmt"

// Error QVM 后端标准错误
type Error struct {
	Code      int    `json:"code"`
	Name      string `json:"name"`
	Message   string `json:"message"`
	RequestID string `json:"request_id"`
}

func (e *Error) Error() string {
	return fmt.Sprintf("reqID:%s code:%d name:%s message:%s", e.RequestID, e.Code, e.Name, e.Message)
}
