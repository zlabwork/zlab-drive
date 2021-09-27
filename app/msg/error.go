package msg

const (
	OK        = 200
	Err       = 400
	ErrNoData = 404
	ErrDB     = 405
)

var statusText = map[int]string{
	OK:        "success",
	Err:       "error",
	ErrDB:     "database error",
	ErrNoData: "can not find",
}

func Text(code int) string {
	return statusText[code]
}
