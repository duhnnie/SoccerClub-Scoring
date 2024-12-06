package arrayHelpers

type ReducerFn[T any, U any] func(U, T, int) U

func Reduce[T any, U any](slice []T, reducer ReducerFn[T, U], initialValue U) U {
	var acc U = initialValue

	for index, item := range slice {
		acc = reducer(acc, item, index)
	}

	return acc
}
