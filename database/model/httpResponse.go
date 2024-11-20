package model

const (
	SUCCESS       = 0
	DEFUALT_ERROR = 9999
)

// HTTPResponse - final response to the api consumers
type HTTPResponse struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message,omitempty"`
}

