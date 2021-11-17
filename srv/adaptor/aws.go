package adaptor

type S3Drive struct {
	Name string
}

func NewS3Drive() *S3Drive {
	return &S3Drive{
		Name: "Aws S3 Drive",
	}
}
