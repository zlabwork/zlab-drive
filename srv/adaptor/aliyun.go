package adaptor

type OssDrive struct {
	Name string
}

func NewOssDrive() *OssDrive {
	return &OssDrive{
		Name: "AliYun OSS Drive",
	}
}
