package base

type Response interface {
  GetCode() int
  GetDescription() string
  GetResponseType() interface{}
}

type BaseResponse struct {
  Code int
  Description string
  ResponseType interface{}
}

func (ref *BaseResponse) GetCode() int {
  return ref.Code
}
func (ref *BaseResponse) GetDescription() string {
  return ref.Description
}
func (ref *BaseResponse) GetResponseType() interface{} {
  return ref.ResponseType
}
