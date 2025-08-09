package response

type Status string

const (
	StatusSuccess Status = "success"
	StatusFailed  Status = "failed"
	StatusError   Status = "error"
)

type Response[T any] struct {
	Status  Status      `json:"status"`
	Message string      `json:"message"`
	Data    *T          `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}
