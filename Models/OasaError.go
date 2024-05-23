package models

const (
	INTERNALL_SERVER_ERROR = "INTERNAL SERVER ERROR"
)

type OasaError struct {
	Error string `json:"error" `
}
