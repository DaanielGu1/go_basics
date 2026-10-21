package tools

import "Treino-em-go/models"

// ---- Interface used for comparation of different types ---- \\

type Comparable[T any] interface {
	equals(other T) bool
}

// ---- Type User ---- \\

type UserComparable struct {
	user *models.User
}

func NewUserComparable(userCmp *models.User) *UserComparable {
	return &UserComparable{user: userCmp}
}

func (cmp *UserComparable) equals(userCmp *models.User) bool {
	return cmp.user == userCmp
}

// ---- Type Group ---- \\

type GroupComparable struct {
	group *models.Group
}

func NewGroupComparable(groupCmp *models.Group) *GroupComparable {
	return &GroupComparable{group: groupCmp}
}

func (cmp *GroupComparable) equals(groupCmp *models.Group) bool {
	return cmp.group == groupCmp
}

// ---- Searcher generic implementations ---- \\

// 1. First Searcher \\

type FindFirstSearcher[T any] struct {
	flag bool
	index int
	obj Comparable[T]
}

func (cmp *FindFirstSearcher[T]) equals(other T) bool {
	return cmp.obj.equals(other)
}

func NewFindSearcher[T any](comparator Comparable[T]) *FindFirstSearcher[T] {
	return &FindFirstSearcher[T]{flag: false, obj: comparator}
}

func (searcher *FindFirstSearcher[T]) Search(index int, other T) {
	if searcher.obj.equals(other) {
		searcher.flag = true
		searcher.index = index
	}
}

func (searcher *FindFirstSearcher[T]) Check(other T) bool {
	return searcher.obj.equals(other)
}

// 2. All Searchers \\

type FindAllSearchers[T any] struct {
	flag bool
	index []int
	obj Comparable[T]
}

func NewFindAll[T any](comparator Comparable[T]) *FindAllSearchers[T] {
	return &FindAllSearchers[T] {
		flag: false,
		obj: comparator,
		index: []int{},
	}
}

func (searcher *FindAllSearchers[T]) Search(index int, other T) {
	if searcher.obj.equals(other) {
		searcher.flag = true
		searcher.index = append(searcher.index, index)
	}
}

func (searcher *FindAllSearchers[T]) Check(other T) bool {
	return searcher.obj.equals(other)
}