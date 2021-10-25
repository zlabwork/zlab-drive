package msg

const (
	OK           = 200
	Err          = 400
	ErrNoData    = 404
	ErrDB        = 405
	ErrParameter = 406
)

var statusText = map[int]string{
	OK:           "success",
	Err:          "error",
	ErrDB:        "database error",
	ErrNoData:    "can not find",
	ErrParameter: "error Parameter",
}

func Text(code int) string {
	return statusText[code]
}
