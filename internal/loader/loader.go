package loader

type Loader[T any] interface {
	Load() (data T, err error)
}
