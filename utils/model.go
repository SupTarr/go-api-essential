package utils

type status string

const (
	Success      status = "SUCCESS"
	Fail         status = "FAILED"
	DataNotFound status = "DATA_NOT_FOUND"
)

type Response struct {
	Status  status `json:"status"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}
