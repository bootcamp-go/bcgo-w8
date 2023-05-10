package api

const (
	// ResponseStatusSuccess is the status of a successful response.
	ResponseStatusSuccess = "success"

	// ResponseStatusFail is the status of a failed response.
	ResponseStatusFail = "fail"

	// ResponseStatusError is the status of an erroneous response.
	ResponseStatusError = "error"
)

// Response is a structure that represents a generic response from the API, similar to JSend format.
type Response struct {
	// Data returned by the API.
	Data any `json:"data,omitempty"`

	// Status is a string that indicates the status of the response.
	// It can be "success", "fail" or "error".
	Status string `json:"status,omitempty"`

	// Message is a string that contains a message about the response.
	Message string `json:"message,omitempty"`
}

// NewSuccessResponse returns a new successful response.
func NewSuccessResponse(data any) Response {
	return Response{
		Data:   data,
		Status: ResponseStatusSuccess,
	}
}

// NewFailResponse returns a new failed response.
func NewFailResponse(data any) Response {
	return Response{
		Data:   data,
		Status: ResponseStatusFail,
	}
}

// NewErrorResponse returns a new erroneous response.
func NewErrorResponse(message string) Response {
	return Response{
		Status:  ResponseStatusError,
		Message: message,
	}
}
