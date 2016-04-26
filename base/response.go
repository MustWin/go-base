package base

// Response describes the minimal interface of web responses.
type Response interface {
	GetCode() int
	GetDescription() string
	GetResponseType() interface{}
}

// DefaultResponse implements Response interface.
type DefaultResponse struct {
	Code         int
	Description  string
	ResponseType interface{}
}

// GetCode returns the HTTP response code.
func (ref *DefaultResponse) GetCode() int {
	return ref.Code
}

// GetDescription returns the text description of the response code.
func (ref *DefaultResponse) GetDescription() string {
	return ref.Description
}

// GetResponseType returns a generic interface{} corresponding to
// the response type.
func (ref *DefaultResponse) GetResponseType() interface{} {
	return ref.ResponseType
}
