package tools

type ISearcher[T any] interface {
	Search(data T, index int)
	EarlyStop() bool
}
type ISimpleSearcher[T any] interface {
	Check(data T) bool
}