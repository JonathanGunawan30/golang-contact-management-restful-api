package http

type ErrorResponse struct {
	Errors string `json:"errors"`
}

type DataEnvelope[T any] struct {
	Data T `json:"data"`
}
