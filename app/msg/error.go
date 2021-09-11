package msg

const (
	OK        = 200
	Err       = 400
	ErrNoData = 404
)

var statusText = map[int]string{
	OK:        "success",
	Err:       "error",
	ErrNoData: "can not find",
}

func Text(code int) string {
	return statusText[code]
}
