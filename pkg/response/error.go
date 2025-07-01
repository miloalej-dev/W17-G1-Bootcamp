package response

type ServiceError struct {
	Code    int    `json:"Code"`
	Message string `json:"Message"`
}

func (e ServiceError) Error() string {
	return e.Message
}

var (
	ErrNotFound       = ServiceError{Code: 404, Message: "document not found"}
	ErrNotImplemented = ServiceError{Code: 501, Message: "not implemented yet!"}
	ErrBadRequest     = ServiceError{Code: 400, Message: "bad request"}
	ErrInternalServer = ServiceError{Code: 500, Message: "internal Server Error"}
	ErrUnauthorized   = ServiceError{Code: 401, Message: "unauthorize token"}
	ErrBadGateway     = ServiceError{Code: 502, Message: "bad gateway"}
)
