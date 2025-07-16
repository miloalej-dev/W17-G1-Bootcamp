package json

type File struct {
	path string
}

func NewFile(path string) *File {
	return &File{
		path: path,
	}
}
