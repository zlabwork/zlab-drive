package adaptor

type LocalDrive struct {
	Name string
}

func NewLocalDrive() *LocalDrive {
	return &LocalDrive{
		Name: "Local File Drive",
	}
}
