package tools

type ISearcher[T any] interface {
	Search(data T, index int)
	EarlyStop() bool
}