package api

// Response an API response
type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data,omitempty"`  // can be an object, a map, or a slice
	Error   string      `json:"error,omitempty"` // probably a string
}

// Success create an api success response.
func Success(data interface{}) Response {
	return Response{
		Success: true,
		Data:    data,
	}
}

// Failed create an api failed response.
func Failed(msg string) Response {
	return Response{
		Error: msg,
	}
}
