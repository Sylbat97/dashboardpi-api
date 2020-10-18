package handler

//BaseHTTPError is an interface for errors to send back to the client
type BaseHTTPError interface {
	Error() string
	// ResponseBody returns response body.
	ResponseBody() ([]byte, error)
	// ResponseHeaders returns http status code and headers.
	ResponseHeaders() (int, map[string]string)
}
