package response

import "errors"

var (
	OK            = Error{Code: 0, Msg: "OK"}
	BadRequest    = Error{Code: 400, Msg: "Bad Request"}
	Unauthorized  = Error{Code: 401, Msg: "Unauthorized"}
	Forbidden     = Error{Code: 403, Msg: "Forbidden"}
	NotFound      = Error{Code: 404, Msg: "Not found"}
	InternalError = Error{Code: 500, Msg: "Internal Error"}
	CanNotDoThis  = errors.New("you can not do this")
)
