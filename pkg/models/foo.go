package models

// Foo is a struct that represents a foo

type Foo struct {
	ID   int
	Name string
}

type FooDoc struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
