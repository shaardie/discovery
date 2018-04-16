package utils

type Test interface {
	Description() string
	Run() (string, error)
}
